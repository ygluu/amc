package amc

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"reflect"
	"time"
)

type iThread interface {
	Start(thread iThread)
	Stop()
	Run()
	Exec(run func())
	IsStart() bool
	IsStop() bool
	IsStoped() bool
	GetWatchMsg() string
	WaitStarted()
	WaitStoped()
}

type thread struct {
	state  int
	isStop bool
}

func (this *thread) GetWatchMsg() string {
	return ""
}

func (this *thread) Run() {
}

func (this *thread) Exec(run func()) {
	for {
		run()
		if this.IsStop() {
			break
		}
	}
}

func (this *thread) doStart(thread iThread) {
	LogI("线程就绪：%s", reflect.TypeOf(thread).Elem().String())

	this.state = 1
	time.Sleep(time.Duration(time.Millisecond * 10))

	thread.Exec(thread.Run)

	LogI("线程结束：%s", reflect.TypeOf(thread).Elem().String())

	this.state = 2
	time.Sleep(time.Duration(time.Millisecond * 10))
}

func (this *thread) Start(thread iThread) {
	go this.doStart(thread)
	for {
		time.Sleep(time.Duration(time.Millisecond * 10))
		if this.state != 0 {
			break
		}
	}
}

func (this *thread) Stop() {
	this.isStop = true
}

func (this *thread) WaitStarted() {
	for !this.IsStart() {
		time.Sleep(time.Millisecond * 10)
	}
}

func (this *thread) WaitStoped() {
	this.Stop()
	for !this.IsStoped() {
		time.Sleep(time.Millisecond * 10)
	}
}

func (this *thread) IsStart() bool {
	return this.state == 1
}

func (this *thread) IsStop() bool {
	return this.isStop
}

func (this *thread) IsStoped() bool {
	return this.state == 2
}

var threads = make([]iThread, 0)

func AddThread(thread iThread) {
	for _, t := range threads {
		if t == thread {
			return
		}
	}
	threads = append(threads, thread)
}

type watch struct {
	thread
}

func (this *watch) outLog(msg string) {
	onWatchLog(msg)
	LogE(msg)
}

func (this *watch) Run() {
	time.Sleep(time.Duration(time.Second * 1))

	msg := GetWatchMsg(&main_svc.ModuleB)
	if msg != "" {
		this.outLog("[amc]WatchThread => " + msg)
	}

	for _, thread := range threads {
		msg := thread.GetWatchMsg()
		if msg != "" {
			this.outLog("[amc]WatchThread => " + msg)
		}
	}
}
