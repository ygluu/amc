package amc

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"runtime"
	"time"
)

type mainSvc struct {
	ModuleST
}

func (this *mainSvc) Queue() *RecvQueue {
	return recv_queue
}

func (this *mainSvc) Init() {
}

func startThreads() {
	for _, t := range threads {
		t.Start(t)
	}
}

func stopThreads() {
	closeNets()

	for _, t := range threads {
		t.Stop()
	}

	for _, t := range threads {
		t.WaitStoped()
	}
}

func (this *mainSvc) Exec(run func()) {
	runtime.LockOSThread()

	registerModules()

	for _, m := range modulesOfCore {
		m.Init()
	}
	for _, m := range modulesOfCore {
		m.Ready()
	}

	for _, m := range modulesOfNormal {
		m.Init()
	}
	for _, m := range modulesOfNormal {
		m.Ready()
	}

	startThreads()

	watch := &watch{}
	watch.Start(watch)
	watch.WaitStarted()

	tevent := &timeEvent{}
	tevent.Start(tevent)
	tevent.WaitStarted()

	runMethodInfos := methodInfosOfId[MsgIdOfMsg(SysMsg_Run)]

	LogI("服务就绪")

	for {
		CurTime = time.Now()
		CurTick = CurTime.UnixMilli()

		procMsgQueue(&main_svc.ModuleB, recv_queue)

		// 发送跑圈消息
		for _, info := range runMethodInfos {
			callMethod(&main_svc.ModuleB, NullCtx, SysMsg_Run, info)
		}

		newTime := time.Now()
		if (recv_queue.GetCount() == 0) && (newTime.UnixMilli()-CurTick < idleRunCheckTime) {
			// 空闲时休眠
			idleSleep()
		}

		if this.IsStop() {
			break
		}
	}

	// 先停止主动线程
	tevent.Stop()
	tevent.WaitStoped()
	watch.Stop()
	watch.WaitStoped()

	stopThreads()

	for i := len(modulesOfNormal) - 1; i >= 0; i-- {
		modulesOfNormal[i].Final()
	}

	for i := len(modulesOfCore) - 1; i >= 0; i-- {
		modulesOfCore[i].Final()
	}
}

func (this *mainSvc) Name() string {
	return "主线程"
}
