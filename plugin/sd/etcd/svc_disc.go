package sd

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

// 服务发现插件

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	etcd3 "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"golang.org/x/net/context"

	"lib/amc"
)

type disc struct {
	clusterName string
	serviceName string
	serviceAddr string
	etcdAddrs   string
	ttl         int
	interval    int
	etcdCfg     *etcd3.Config
	etcdCli     *etcd3.Client
	weight      uint32
}

func NewEtcdCfg(endpoints string) *etcd3.Config {
	ret := &etcd3.Config{}
	ret.Endpoints = strings.Split(endpoints, ";")
	return ret
}

func (this *disc) conn() *etcd3.Client {
	if this.etcdCfg.DialTimeout == 0 {
		this.etcdCfg.DialTimeout = 3 * time.Second
	}

	client, err := etcd3.New(*this.etcdCfg)
	if err != nil {
		amc.LogE("连接ETCD失败: %s", err.Error())
		return nil
	}

	amc.LogI("连接ETCD成功: %s", this.etcdAddrs)
	return client
}

const addrskeyname = "svcaddrs"
const idskeyname = "msgids"

func (this *disc) getSvcRootKey() string {
	return strings.ToLower(fmt.Sprintf("/%s/%s/", this.clusterName, addrskeyname))
}

func (this *disc) getSvcAddrKey() string {
	return this.getSvcRootKey() + strings.ToLower(fmt.Sprintf("%s/%s",
		this.serviceName, this.serviceAddr))
}

func (this *disc) getSvcMsgIdsKey() string {
	return strings.ToLower(fmt.Sprintf("/%s/%s/%s/%s", this.clusterName,
		idskeyname, this.serviceName, this.serviceAddr))
}

func toJsonStr(v interface{}) string {
	ret, _ := json.Marshal(v)
	return string(ret)
}

func (this *disc) MsgIds(etcdKey string) string {
	if this.etcdCli == nil {
		return ""
	}

	etcdKey = strings.Replace(etcdKey, addrskeyname, idskeyname, 1)

	resp, err := this.etcdCli.Get(context.Background(), etcdKey, etcd3.WithPrefix())
	if err != nil {
		return ""
	}

	if resp != nil && resp.Kvs != nil {
		for _, kv := range resp.Kvs {
			return string(kv.Value)
		}
	}

	return ""
}

func (this *disc) registerSvc(addr string, msgids []uint32) {
	isFirstLog := false

	checkErrr := func(key string, err error) {
		if err == nil {
			if isFirstLog {
				isFirstLog = false
				amc.LogI("服务注册成功：%s", key)
			}
		} else {
			amc.LogE("服务注册失败：%s, %s", err, key)
		}
	}

	doRegister := func(key, value string, ttl int) {
		resp, err := this.etcdCli.Grant(context.TODO(), int64(ttl))
		if err != nil {
			return
		}

		_, err = this.etcdCli.Get(context.Background(), key)

		if err == nil {
			_, err := this.etcdCli.Put(context.Background(), key, value, etcd3.WithLease(resp.ID))
			checkErrr(key, err)
		} else if err == rpctypes.ErrKeyNotFound {
			_, err := this.etcdCli.Put(context.TODO(), key, value, etcd3.WithLease(resp.ID))
			checkErrr(key, err)
		} else {
			amc.LogE("获取服键值失败：%s, %s", err, key)
		}
	}

	timeCount := 0

	for {
		time.Sleep(time.Second)
		if this.etcdCli == nil {
			this.etcdCli = this.conn()
			if this.etcdCli == nil {
				continue
			}

			doRegister(this.getSvcMsgIdsKey(), toJsonStr(msgids), 0xFFFFFFFF)
			isFirstLog = true
			// 启动注册：1-weight
			doRegister(this.getSvcAddrKey(), fmt.Sprintf("1-%d", this.weight), this.ttl)
			timeCount = 0
		}

		timeCount++
		if timeCount < this.interval {
			continue
		}
		timeCount = 0
		// 续约注册：2-weight
		doRegister(this.getSvcAddrKey(), fmt.Sprintf("2-%d", this.weight), this.ttl)
	}
}

func (this *disc) ParseName(key string) (svcname, addr string) {
	arr := strings.Split(key, "/")
	len := len(arr)
	if len < 2 {
		return
	}

	return arr[len-2], arr[len-1]
}

func (this *disc) Watch(onWatch func(key string, weight, flag int)) {
	// flag: 0：删除，1：上线，2：续约
	go this.watch(onWatch)
}

func (this *disc) watch(onWatch func(key string, weight, flag int)) {
	for {
		if this.etcdCli != nil {
			break
		}
		time.Sleep(time.Second)
	}

	procWatch := func(key, value string, isDel bool) {
		if isDel {
			onWatch(key, 0, 0)
			return
		}

		vs := strings.Split(value, "-")
		if len(vs) != 2 {
			return
		}

		flag, err := strconv.Atoi(vs[0])
		if err != nil {
			return
		}

		weight, err := strconv.Atoi(vs[1])
		if err != nil {
			return
		}

		onWatch(key, weight, flag)
	}

	client := this.etcdCli
	rootKey := this.getSvcRootKey()

	for {
		resp, err := client.Get(context.Background(), rootKey, etcd3.WithPrefix())
		if err != nil {
			time.Sleep(time.Second)
			continue
		}

		if resp != nil && resp.Kvs != nil {
			for _, kv := range resp.Kvs {
				procWatch(string(kv.Key), string(kv.Value), false)
			}
		}

		rch := client.Watch(context.Background(), rootKey, etcd3.WithPrefix())
		for wresp := range rch {
			if wresp.Events == nil {
				return
			}
			for _, ev := range wresp.Events {
				switch ev.Type {
				case mvccpb.PUT:
					procWatch(string(ev.Kv.Key), string(ev.Kv.Value), false)
				case mvccpb.DELETE:
					procWatch(string(ev.Kv.Key), "0", true)
				}
			}
		}
	}
}

func New(clusterName, serviceName, serviceAddr, etcdAddrs string, msgids []uint32, etcdCfg *etcd3.Config) amc.ISD {
	ret := &disc{
		clusterName: clusterName,
		serviceName: serviceName,
		serviceAddr: serviceAddr,
		etcdAddrs:   etcdAddrs,
		interval:    20,
		ttl:         25,
		etcdCfg:     etcdCfg,
	}

	if etcdCfg == nil {
		ret.etcdCfg = NewEtcdCfg(etcdAddrs)
	}

	go ret.registerSvc(serviceAddr, msgids)

	return ret
}
