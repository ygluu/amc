package amc

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"reflect"
	"runtime"
)

type Sender struct {
	ModuleST
	msgQueue *SendQueue
	net      INet
	lb       ILB
	codec    ICodec
}

func NewSender(msgQueue *SendQueue, net INet, codec ICodec, lb ILB) iThread {
	this := &Sender{}
	this.net = net
	this.msgQueue = msgQueue
	this.codec = codec
	this.lb = lb

	if LB == nil {
		LB = lb
	}

	if send_queue == nil {
		send_queue = msgQueue
	}

	RegisterModuleIns(nil, this)
	if lb != nil {
		RegisterModuleIns(this, lb.(IModule))
	}

	return this
}

func (this *Sender) Name() string {
	return "发送线程"
}

func (this *Sender) Exec(run func()) {
	// 绑定线程，系统级调度
	runtime.LockOSThread()
	this.ModuleST.Exec(run)
}

func (this *Sender) Run() {
	this.ModuleST.Run()

	onMsg := func(msgNode *SendNode) {
		ctx := msgNode.Ctx
		msg := msgNode.Msg

		if (ctx == nil) || (msg == nil) {
			this.isStop = true
			return
		}

		msgid, isProto := MsgInfoOfMsg(msg)
		if !isProto {
			infos := methodInfosOfId[msgid]
			for _, info := range infos {
				info.msgQueue.Push(ctx, msg, info)
			}
			return
		}

		msgBuf, err := this.codec.Marshal(msg)
		if err != nil {
			typ := reflect.TypeOf(msg)
			LogE("[amc]Sender.run => 消息序列化失败<%s>", typ.Elem().Name())
			return
		}

		conn := ctx.GetConn()
		if conn != nil {
			err = this.net.SendPack(conn, msgid, ctx.GetFlag(), msgBuf)
			if err != nil {
				LogE("[amc]Sender.run => 发送失败<%s>", err.Error())
			}
			return
		}

		flag := ctx.GetFlag()
		conns := this.lb.GetConns(msgid)
		for _, conn := range conns {
			this.net.SendPack(conn, msgid, flag, msgBuf)
		}
	}

	nodes := this.msgQueue.Pop()
	for {
		node := nodes
		if node == nil {
			return
		}
		nodes = nodes.Next
		onMsg(node)
	}

	if this.msgQueue.GetCount() == 0 {
		idleSleep()
	}
}
