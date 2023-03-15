package svc

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
	"lib/amc/plugin/disc/etcd"
	// 网络通信插件
	"lib/amc/plugin/tcp"
)

func Exec() {
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
						|  		 (内部消息发送线程)  (内部消息发送队列)     (网关线程)           |                          |
						|	                                                        	 |                          |
						|	                                                        	 |                          |
						|										      (内部服务的消息)|-->--|                          |
						|	                                                       | 	                            |
						|	                                                       | 	                            |
						|----DataPack------>------------------------------------[CSender]----->-----DataPack--------|
						                                        				   (客户端发送线程)


	******************************************************************************************************************************************************/

	var listenAddr string
	var etcdAddrs string
	var clusterName string
	var svcName string

	flag.StringVar(&clusterName, "CName", "异步微服务集群", "ClusterName(集群名称)，默认：异步微服务集群")
	flag.StringVar(&svcName, "SName", "微服务网关", "ServiceName(服务名称)，默认：微服务网关")
	flag.StringVar(&listenAddr, "LAddr", "127.0.0.1:4330", "ListenAddr(外网监听地址)，默认：127.0.0.1:4330")
	flag.StringVar(&etcdAddrs, "EAddrs", "127.0.0.1:2379;127.0.0.1:2379", "ETCD Cluster Addr List(集群地址列表)，默认：127.0.0.1:2379;127.0.0.1:2379")

	// 链接内部集群的网络插件和内部通信封包器
	snet := tcp.New(scodep.New())
	// 链接客户端的网络插件和客户端通信封包器
	cnet := tcp.New(ccodep.New())

	// 消息接收队列
	RMsgQueue := amc.NewQueue()
	// 消息发送队列
	SMsgQueue := amc.NewQueue()

	// ProtoBuf编码器（序列化反序列化插件）
	codec := pb.New()
	// 服务发现插件
	disc := disc.New(etcdAddrs, snet)

	// 客户端数据接收线程
	CRecver := amc.NewGateRecver(RMsgQueue, snet, cnet, codec, disc)
	// 客户端数据发送线程
	CSender := amc.NewGateSender(RMsgQueue, snet, cnet, codec, disc)
	// 内部集群消息发送线程
	SSender := amc.NewSender(SMsgQueue, snet, codec, disc)

	// 对象统一管理，进程结束时先关闭网络再终止线程
	amc.AddNet(snet)
	amc.AddNet(cnet)
	amc.AddThread(CRecver)
	amc.AddThread(CSender)
	amc.AddThread(SSender)

	// 监听客户端
	err := cnet.ListenByAddr(listenAddr)
	if err != nil {
		amc.LogE("监听失败：" + listenAddr)
		return
	}

	// 服务器启动运行
	amc.Exec(clusterName, svcName, listenAddr, snet, RMsgQueue, RMsgQueue)
}
