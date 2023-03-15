package amc

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"sync"
)

type msgQueue struct {
	mutex sync.Mutex
	count uint
}

func (this *msgQueue) GetCount() uint {
	return this.count
}

type RecvNode struct {
	Ctx   ICtx
	Msg   IMsg
	minfo *methodInfo
	Next  *RecvNode
}

type RecvQueue struct {
	msgQueue
	head *RecvNode
	tail *RecvNode
}

func NewRecvQueue() *RecvQueue {
	return &RecvQueue{}
}

func (this *RecvQueue) Push(ctx ICtx, msg IMsg, minfo *methodInfo) {
	node := &RecvNode{
		Ctx:   ctx,
		Msg:   msg,
		minfo: minfo,
	}

	this.mutex.Lock()

	if (this.head == nil) || (this.tail == nil) {
		this.head = node
		this.tail = node
	} else {
		this.tail.Next = node
		this.tail = node
	}
	this.count++

	this.mutex.Unlock()
}

func (this *RecvQueue) Pop() (ret *RecvNode) {

	this.mutex.Lock()

	ret = this.head
	this.head = nil
	this.tail = nil
	this.count = 0

	this.mutex.Unlock()

	return ret
}

type SendNode struct {
	Ctx  ICtx
	Msg  IMsg
	Next *SendNode
}

type SendQueue struct {
	msgQueue
	head *SendNode
	tail *SendNode
}

func NewSendQueue() *SendQueue {
	return &SendQueue{}
}

func (this *SendQueue) Push(ctx ICtx, msg IMsg) {
	node := &SendNode{
		Ctx: ctx,
		Msg: msg,
	}

	this.mutex.Lock()

	if (this.head == nil) || (this.tail == nil) {
		this.head = node
		this.tail = node
	} else {
		this.tail.Next = node
		this.tail = node
	}
	this.count++

	this.mutex.Unlock()
}

func (this *SendQueue) Pop() (ret *SendNode) {
	this.mutex.Lock()

	ret = this.head
	this.head = nil
	this.tail = nil
	this.count = 0

	this.mutex.Unlock()

	return ret
}
