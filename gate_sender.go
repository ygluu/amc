package amc

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

type GateSender struct {
	Recver
	snet INet
	cnet INet
}

func NewGateSender(msgQueue *RecvQueue, snet, cnet INet) iThread {
	this := &GateSender{}
	this.snet = snet
	this.cnet = cnet
	this.net = snet
	if recv_queue == nil {
		recv_queue = msgQueue
	}

	return this
}

func (this *GateSender) run() {
	conn, addr, msgid, flag, pack := this.snet.RecvPack()
	if (conn == nil) || (msgid == 0) || (pack == nil) {
		return
	}

	// 是发给服务网关的
	if isMyMsg(msgid) {
		this.UnmarshalAndEnqueue(conn, addr, msgid, flag, pack)
		return
	}

	if flag == 0 {
		LogW("[amc]GateSender.run => 消息发到无效客户端链接：%d", msgid)
		return
	}

	// 是发个客户端的
	err := this.cnet.SendPack(IConn(nil), msgid, 0, pack)
	if err != nil {
		LogE("[amc]GateSender.run => 发送消息错误：%v", err)
	}
}
