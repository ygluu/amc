package amc

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"runtime"
)

type Recver struct {
	ModuleST
	net   INet
	codec ICodec
}

func NewRecver(msgQueue *RecvQueue, net INet, codec ICodec) iThread {
	this := &Recver{}
	this.net = net
	this.codec = codec
	if recv_queue == nil {
		recv_queue = msgQueue
	}

	return this
}

// 反序列化并加入接收入列
func (this *Recver) UnmarshalAndEnqueue(conn IConn, addr string, msgid uint32, flag uintptr, pack []byte) {
	msg := NewMsgOfId(msgid)
	if msg == nil {
		LogE("[amc]Recver.UnmarshalAndEnqueue => 无效通信命令：%d", msgid)
		return
	}

	infos := methodInfosOfId[msgid]
	if infos == nil {
		LogE("[amc]Recver.UnmarshalAndEnqueue => 无主消息：%d", msgid)
		return
	}

	err := this.codec.Unmarshal(pack, msg)
	if err != nil {
		LogE("[amc]Recver.UnmarshalAndEnqueue => 反序列化失败：%v", err)
		return
	}

	ctx := &recvCtx{
		conn: conn,
		flag: flag,
		addr: addr,
	}

	for _, info := range infos {
		info.msgQueue.Push(ctx, msg, info)
	}
}

func (this *Recver) Run() {
	conn, addr, msgid, flag, pack := this.net.RecvPack()
	if (conn == nil) || (msgid == 0) || (pack == nil) {
		return
	}

	this.UnmarshalAndEnqueue(conn, addr, msgid, flag, pack)
}

func (this *Recver) Exec(run func()) {
	// 绑定线程，系统级调度
	runtime.LockOSThread()
	this.ModuleST.Exec(run)
}
