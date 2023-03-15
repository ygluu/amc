package main

import (
	"flag"

	"lib/amc"

	// ProtoBuf编码器插件
	"lib/amc/plugin/codec/pb"
	// 集群内部通信数据封包器插件
	"lib/amc/plugin/codep/client"

	// 负载均衡插件
	"lib/amc/plugin/lb"

	// 网络通信插件
	"lib/amc/plugin/tcp"

	"lib/amc/demo/client/test"
)

func main() {
	clusterName := "异步微服务集群"
	serviceName := "客户端"
	hostName := "压力测试"
	flag.StringVar(&clusterName, "CName", clusterName, "集群名称，默认："+clusterName)
	flag.StringVar(&serviceName, "SName", serviceName, "服务名称，默认："+serviceName)

	gateAddrs := "127.0.0.1:1921;127.0.0.1:1921"
	flag.StringVar(&gateAddrs, "GateAddrs", gateAddrs, "微服务网关地址列表，默认："+gateAddrs)

	hb := lb.NewHashBalan(1000)
	hb.AddAddrs(serviceName, gateAddrs)
	hb.Calc()

	// 网络插件
	net := tcp.New(ccodep.New())

	// 主要是日志输出带有相关信息
	amc.SetSvc(clusterName, serviceName, hostName)

	// 消息接收队列
	RQueue := amc.NewRecvQueue()
	// 消息发送队列
	SQueue := amc.NewSendQueue()

	// 序列化反序列化插件（ProtoBuf编码器）
	codec := pb.New()

	// 接收线程
	Recver := amc.NewRecver(RQueue, net, codec)
	// 发送线程
	Sender := amc.NewSender(SQueue, net, codec, nil)

	// 对象统一管理，进程结束时先关闭网络再终止线程
	amc.AddNet(net)
	amc.AddThread(Recver)
	amc.AddThread(Sender)

	test.Start(net, hb, codec)

	// 服务器启动运行
	amc.Exec()
}
