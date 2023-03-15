package lb

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"math"
	"sync"

	"lib/amc"
)

type msgSvcWatch struct {
	amc.Msg
	key    string
	weight int
	flag   int
}

type msgAddConn struct {
	amc.Msg
	conn amc.IConn
	addr string
}

type SvcInfo struct {
	hash     uint32
	name     string
	addr     string
	hitCount int
	weight   int
	msgids   []uint32
	conn     amc.IConn
}

func (this *SvcInfo) Name() string {
	return this.name
}

func (this *SvcInfo) Addr() string {
	return this.addr
}

// 每个loadBalan只能有一个宿主线程模块来处理消息和调用loadBalan方法
//    默认发送线程Sender和GateRevcer自动绑定loadBalan，参阅amc.RegisterModuleIns）
type loadBalan struct {
	amc.ModuleB

	mutex      sync.Mutex
	sd         amc.ISD
	net        amc.INet
	replicas   int
	weightGran int

	// svcInfoOfAddr 仅为loadBalan的宿主线程使用，所以不用锁
	svcInfoOfAddr map[string]*SvcInfo

	// 以下为多线程接口使用，用时需要加锁
	hashBalanOfMsgId map[uint32]*HashBalan
}

func (this *loadBalan) Name() string {
	return "负载均衡"
}

func (this *loadBalan) GetHash(key string) uint32 {
	return getHash(key)
}

func (this *loadBalan) getConn(info *SvcInfo) (ret amc.IConn) {
	ret = info.conn
	if (ret == nil) || ret.IsClose() {
		ret, _ = this.net.Dial(info.addr)
		info.conn = ret
	}
	return ret
}

func (this *loadBalan) GetConn(msgid uint32) amc.IConn {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	HashBalan := this.hashBalanOfMsgId[msgid]
	if HashBalan == nil {
		return nil
	}

	info := HashBalan.Get()
	if info == nil {
		return nil
	}

	return this.getConn(info)
}

func (this *loadBalan) GetConns(msgid uint32) []amc.IConn {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	HashBalan := this.hashBalanOfMsgId[msgid]
	if HashBalan == nil {
		return []amc.IConn{}
	}

	ret := make([]amc.IConn, HashBalan.infoCount)

	count := 0
	for _, node := range HashBalan.hashNodes {
		for _, info := range node.infos {
			conn := this.getConn(info)
			if conn == nil {
				continue
			}
			ret[count] = conn
			count++
		}
	}

	return ret[:count]
}

func (this *loadBalan) CheckGetConn(msgid uint32, checkSvcName string, DefConn amc.IConn) amc.IConn {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	HashBalan := this.hashBalanOfMsgId[msgid]
	if HashBalan == nil {
		return nil
	}

	info := HashBalan.Get()
	if info == nil {
		return nil
	}

	if info.name == checkSvcName {
		return DefConn
	}

	return this.getConn(info)
}

func (this *loadBalan) getHash(key string) uint32 {
	data := md5.Sum([]byte(key))
	key = fmt.Sprintf("%x", data)
	return crc32.ChecksumIEEE([]byte(key))
}

func (this *loadBalan) GetConnByKey(msgid uint32, key string) amc.IConn {
	return this.GetConnByHash(msgid, this.getHash(key))
}

func (this *loadBalan) GetConnByHash(msgid, hash uint32) amc.IConn {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	hb := this.hashBalanOfMsgId[msgid]
	if hb == nil {
		return nil
	}

	info := hb.GetByHash(hash)
	if info == nil {
		return nil
	}

	return this.getConn(info)
}

func (this *loadBalan) GetConnByAddr(addr string) amc.IConn {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	info := this.svcInfoOfAddr[addr]
	if info == nil {
		return nil
	}
	return this.getConn(info)
}

func (this *loadBalan) GetInfo(msgid uint32) (name, addr string) {
	return this.GetInfoByHash(msgid, 0)
}

func (this *loadBalan) GetInfoByKey(msgid uint32, key string) (name, addr string) {
	return this.GetInfoByHash(msgid, getHash(key))
}

func (this *loadBalan) GetInfoByHash(msgid, hash uint32) (name, addr string) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	hb := this.hashBalanOfMsgId[msgid]
	if hb == nil {
		return "", ""
	}

	var info *SvcInfo

	if hash == 0 {
		info = hb.Get()
	} else {
		info = hb.GetByHash(hash)
	}
	if info == nil {
		return "", ""
	}

	return info.name, info.addr
}

func (this *loadBalan) AddConn(addr string, conn amc.IConn) {
	msg := &msgAddConn{
		conn: conn,
		addr: addr,
	}
	this.SendMsg(amc.NullCtx, msg)
}

func (this *loadBalan) OnMsg_AddConn(ctx amc.ICtx, msg *msgAddConn) {
	info := this.svcInfoOfAddr[msg.addr]
	if info == nil {
		return
	}

	if info.conn != nil {
		return
	}

	this.mutex.Lock()
	defer this.mutex.Unlock()
	info.conn = msg.conn
}

func (this *loadBalan) recalc() {
	hashBalanOfMsgId := make(map[uint32]*HashBalan)

	// 先按消息ID对服务进行分类
	for _, info := range this.svcInfoOfAddr {
		for _, msgid := range info.msgids {
			hb := hashBalanOfMsgId[msgid]
			if hb == nil {
				hb = NewHashBalan(this.replicas)
				hashBalanOfMsgId[msgid] = hb
			}
			hb.Add(info)
		}
	}

	// 每个消息ID的服务列表进行哈希散列计算
	for _, hb := range hashBalanOfMsgId {
		hb.Calc()
	}

	// 加锁赋值
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.hashBalanOfMsgId = hashBalanOfMsgId
}

func (this *loadBalan) OnMsg_SvcWatch(ctx amc.ICtx, msg *msgSvcWatch) {
	key := msg.key
	weight := msg.weight
	flag := msg.flag // flag: 0：删除，1：上线，2：续约

	name, addr := this.sd.ParseName(key)

	if flag == 0 {
		amc.LogI("[amc]loadBalan.OnMsg_SvcWatch => 服务下线，name:%s, addr:%s", name, addr)
		this.recalc()
		return
	}

	info := this.svcInfoOfAddr[addr]
	isExist := info != nil

	getMsgIds := func() (ret []uint32) {
		// 从ETCD获取该服务方法对应的消息ID列表
		msgids := this.sd.MsgIds(key)
		json.Unmarshal([]byte(msgids), &ret)
		return
	}

	if !isExist {
		info = &SvcInfo{
			name:   name,
			addr:   addr,
			weight: weight,
			msgids: getMsgIds(),
		}
		this.svcInfoOfAddr[addr] = info
		amc.LogI("[amc]loadBalan.OnMsg_SvcWatch => 服务上线，name:%s, addr:%s, weight:%d, flag:%d, MsgIds:%v",
			name, addr, weight, flag, info.msgids)
	} else if info.name != name {
		amc.LogI("[amc]loadBalan.OnMsg_SvcWatch => 服务地址重复，name1:%s, name2:%s, addr:%s",
			name, info.name, addr)
		return
	} else if flag == 1 {
		// 如果已经存在，但服务重启上线，更新消息ID列表
		info.msgids = getMsgIds()
	}

	// 如果是新上线或者权重抖动超过颗粒度则重新计算哈希分布
	isRecalc := (!isExist) || (flag == 1) || (int(math.Abs(float64(weight-info.weight))) >= this.weightGran)
	if isRecalc {
		this.recalc()
	}
}

// replicas: 每个服务地址的哈希副本数量，数量越大越均衡，0为默认100
//			 权重分配公式：replicas * 服务节点数量 * 该节点权重 / 所有节点权重之和
// weightGran: 权重颗粒度（Weight Granularity），
//			 当节点权重变化大于weightGran时才会重新计算权重均衡分配，0为默认10
func New(sd amc.ISD, net amc.INet, replicas, weightGran int) amc.ILB {
	if replicas <= 0 {
		replicas = 100
	}
	if weightGran <= 0 {
		weightGran = 10
	}

	ret := &loadBalan{
		hashBalanOfMsgId: make(map[uint32]*HashBalan),
		svcInfoOfAddr:    make(map[string]*SvcInfo),
		sd:               sd,
		net:              net,
		weightGran:       weightGran,
		replicas:         replicas,
	}

	// 启动服务发现
	sd.Watch(func(key string, weight, flag int) {
		msg := &msgSvcWatch{
			key:    key,
			weight: weight,
			flag:   flag,
		}
		// 当节点信息发生变化是发送消息给负载均衡器的使用者线程来处理
		ret.SendMsg(amc.NullCtx, msg)
	})

	return ret
}
