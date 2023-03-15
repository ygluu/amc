package main

import (
	"flag"

	"lib/amc"

	// ProtoBuf编码器插件
	"lib/amc/plugin/codec/pb"
	// 客户端通信数据封包器插件
	"lib/amc/plugin/codep/client"
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
	/*****************************************************************************************************************************************************


		                                             							   (客户端接收线程)
		  				<----DataPack----------<--------------------------------[CRecver]----<-----DataPack----------
						|	   			|                                          |                           		|
						|	   			|                                          |                           		|
						|	   			|					    		 (客户端消息)|-->--|                			|
						|	   			|                  	                             |                          |
						|	   			|                  	                             |                          |
						|	            |                                            	 |                          |
	[SideCluster]<--->[snet]   		[SSender]---<---SMsgQueue---<---[SvcThread]---<---RMsgQueue(服务消息接收队列)    [cnet]<--->[RemoteClient](客户端通信)
						|  		 (内部消息发送线程)  (内部消息发送队列)     (网关线程)          |                          |
						|	                                                        	 |                          |
						|	                                                        	 |                          |
						|										      (内部服务的消息)|-->--|                          |
						|	                                                       | 	                            |
						|	                                                       | 	                            |
						|----DataPack------>------------------------------------[CSender]----->-----DataPack--------|
						                                        				   (客户端发送线程)


	******************************************************************************************************************************************************/

	clusterName := "异步微服务集群"
	serviceName := "业务服务"
	flag.StringVar(&clusterName, "CName", clusterName, "集群名称，默认："+clusterName)
	flag.StringVar(&serviceName, "SName", serviceName, "服务名称，默认："+serviceName)

	etcdAddrs := "127.0.0.1:2379;127.0.0.2:2379"
	flag.StringVar(&etcdAddrs, "EtcdAddrs", etcdAddrs, "集群地址列表，默认："+etcdAddrs)

	listenAddr := "127.0.0.1:1921"
	flag.StringVar(&listenAddr, "ListenAddr", listenAddr, "外网监听地址，默认："+listenAddr)

	// 链接客户端的网络插件和客户端通信封包器
	cnet := tcp.New(ccodep.New())
	// 监听客户端
	err := cnet.ListenByAddr(listenAddr)
	if err != nil {
		amc.LogE("监听失败：" + listenAddr)
		return
	}

	// 链接内部集群的网络插件和内部通信封包器
	snet := tcp.New(scodep.New())

	// 主要是日志输出带有相关信息
	amc.SetSvc(clusterName, serviceName, listenAddr)

	// 消息接收队列
	RMsgQueue := amc.NewRecvQueue()
	// 消息发送队列
	SMsgQueue := amc.NewSendQueue()

	// 序列化反序列化插件（ProtoBuf编码器）
	codec := pb.New()
	// 服务发现插件
	sd := sd.New(clusterName, serviceName, listenAddr, etcdAddrs, amc.ProtoMsgIds(), nil)
	// 负载均衡插件
	lb := lb.New(sd, snet, 0, 0)

	// 客户端数据接收线程
	CRecver := amc.NewGateRecver(RMsgQueue, snet, cnet, lb)
	// 客户端数据发送线程
	CSender := amc.NewGateSender(RMsgQueue, snet, cnet)
	// 内部集群消息发送线程
	SSender := amc.NewSender(SMsgQueue, snet, codec, lb)

	// 对象统一管理，进程结束时先关闭网络再终止线程
	amc.AddNet(snet)
	amc.AddNet(cnet)
	amc.AddThread(CRecver)
	amc.AddThread(CSender)
	amc.AddThread(SSender)

	// 服务器启动运行
	amc.Exec()
}
