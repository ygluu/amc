package amc

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var cluster_name = ""
var svc_name = ""
var svc_addr = ""
var LB ILB = nil
var recv_queue = (*RecvQueue)(nil)
var send_queue = (*SendQueue)(nil)
var is_stop = false

var main_svc = &mainSvc{}

func waitStopSig() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	<-ch
	is_stop = true
	time.Sleep(time.Duration(time.Millisecond * 20))
}

func GetClusterName() string {
	return cluster_name
}

func GetSvcName() string {
	return svc_name
}

func SetSvc(clusterName, svcName, svcAddr string) {
	cluster_name = clusterName
	svc_name = svcName
	svc_addr = svcAddr
	LogI("启动中.....")
	LogI("CPU数量：%d", runtime.NumCPU())
	LogI("最大并发：%d", runtime.NumCPU())
}

func Exec() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	main_svc.Start(main_svc)
	main_svc.WaitStarted()

	waitStopSig()

	main_svc.Stop()
	main_svc.WaitStoped()

	LogI("已优雅退出")
	LogI("将在5秒后结束......")
	time.Sleep(5 * time.Second)
}
