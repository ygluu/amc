package main

import (
	"flag"

	"lib/amc"
	_ "lib/amc/demo/proto"

	// ProtoBuf编码器插件
	"lib/amc/plugin/codec/pb"
	// 集群内部通信数据封包器插件
	"lib/amc/plugin/codep/side"

	// etcd服务发现插件
	"lib/amc/plugin/sd/etcd"
	// 负载均衡插件
	"lib/amc/plugin/lb"

	// 网络通信插件
	"lib/amc/plugin/tcp"
)

func main() {
	clusterName := "异步微服务集群"
	serviceName := "业务服务"
	flag.StringVar(&clusterName, "CName", clusterName, "集群名称，默认："+clusterName)
	flag.StringVar(&serviceName, "SName", serviceName, "服务名称，默认："+serviceName)

	etcdAddrs := "127.0.0.1:2379;127.0.0.2:2379"
	flag.StringVar(&etcdAddrs, "EtcdAddrs", etcdAddrs, "集群地址列表，默认："+etcdAddrs)

	ip := amc.GetLoaclIp(true)
	flag.StringVar(&ip, "IP", ip, "服务监听的IP地址，默认：0.0.0.0")

	var beingPort int
	var endPort int
	flag.IntVar(&beingPort, "BPort", 10000, "监听端口起始值，用于随机取监听端口，默认：10000")
	flag.IntVar(&endPort, "EPort", 65535, "监听端口结束值，用于随机取监听端口，默认：65535")

	// 网络插件
	net := tcp.New(scodep.New())

	// 监听客户端
	addr, err := net.ListenByIp(ip, beingPort, endPort)
	if err != nil {
		amc.LogE("监听失败：%s", err.Error())
		return
	}

	// 主要是日志输出带有相关信息
	amc.SetSvc(clusterName, serviceName, addr)

	// 消息接收队列
	RQueue := amc.NewRecvQueue()
	// 消息发送队列
	SQueue := amc.NewSendQueue()

	// 序列化反序列化插件（ProtoBuf编码器）
	codec := pb.New()
	// 服务发现插件
	sd := sd.New(clusterName, serviceName, addr, etcdAddrs, amc.ProtoMsgIds(), nil)
	// 负载均衡插件
	lb := lb.New(sd, net, 0, 0)

	// 接收线程
	Recver := amc.NewRecver(RQueue, net, codec)
	// 发送线程
	Sender := amc.NewSender(SQueue, net, codec, lb)

	// 对象统一管理，进程结束时先关闭网络再终止线程
	amc.AddNet(net)
	amc.AddThread(Recver)
	amc.AddThread(Sender)

	// 服务器启动运行
	amc.Exec()
}
