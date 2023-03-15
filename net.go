package amc

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"container/list"
	"sync"
)

type recvCtx struct {
	conn IConn
	flag uintptr
	addr string
}

func (this *recvCtx) GetConn() IConn {
	return this.conn
}

func (this *recvCtx) GetFlag() uintptr {
	return this.flag
}

func (this *recvCtx) GetAddr() string {
	return this.addr
}

type iTask interface {
	DoExec()
}

type Tasks struct {
	mut  sync.Mutex
	list list.List
}

func (this *Tasks) Add(task iTask) {
	defer this.mut.Unlock()
	this.mut.Lock()
	this.list.PushFront(task)
}

func (this *Tasks) Exec() {
	if this.list.Len() == 0 {
		return
	}

	this.mut.Lock()
	tasks := make([]iTask, this.list.Len())
	index := 0
	for e := this.list.Front(); e != nil; e = e.Next() {
		tasks[index] = e.Value.(iTask)
	}
	this.mut.Unlock()

	for _, task := range tasks {
		task.DoExec()
	}
}

var nets = make([]INet, 0)

func AddNet(net INet) {
	nets = append(nets, net)
}

func closeNets() {
	for _, net := range nets {
		net.Close()
	}
}
