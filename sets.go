package amc

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"fmt"
	"log"
	"net"
	"time"
)

//******************************************************************************
// 写日志函数
var logger = log.Printf

func SetLogger(fn func(format string, v ...any)) {
	logger = fn
}

//******************************************************************************
// DEBUG模式，为false时LogD不输出日志
var isDebugMode = false

func SetDebugMode(value bool) {
	isDebugMode = value
}

//******************************************************************************
// 消息请求频繁超限次数，超过则断线
var cdSpeedingLimit = 10

// 超速超过次数后回调
var onSpeedingLimit = func(addr string, msgid uint32) {}

// 默认消息请求CD时间
var cdTimeDef = int64(1000) // 单位：毫秒

// CD误差值，即可以提交前x毫秒
var cdErrorValue = int64(5) // 单位：毫秒

func SetSpeedingLimit(value int) {
	cdSpeedingLimit = value
}

func SetSpeedingLimitOnFunc(value func(addr string, msgid uint32)) {
	onSpeedingLimit = value
}

func SetCdTimeDef(value int64) {
	cdTimeDef = value
}

func SetCdErrorValue(value int64) {
	cdErrorValue = value
}

//******************************************************************************
// 看守线程检查主线程处理业务超时时间，超过则认为进入死循环
var watchTime = int64(5 * 1000) // 单位：毫秒

func SetWatchTime(value int64) {
	if value == 0 {
		value = 5 * 1000
	}
	watchTime = value
}

// 死循环后回调
var onWatchLog = func(msg string) {}

func SetWatchLogFunc(value func(msg string)) {
	onWatchLog = value
}

//******************************************************************************
// 线程空闲时休眠时间
var idleSleepTime = int64(1) // 单位：毫秒

// 主线程跑圈一次小于这个时间且无消息待处理时则进入休眠状态
var idleRunCheckTime = int64(20) // 单位：毫秒

func SetIdleSleepTime(value int64) {
	idleSleepTime = value
}

func SetIdleRunCheckTime(value int64) {
	idleRunCheckTime = value
}

// 休眠函数，默认调用time.Sleep，也可以根据目标操作系统设置Sleep，单位毫秒
var Sleep = func(MillSec int64) {
	time.Sleep(time.Millisecond * time.Duration(MillSec))
}

func idleSleep() {
	Sleep(idleSleepTime)
}

//******************************************************************************
// 时间跳变消息开关，true时触发时间跳变消息的派发
var switchTimeMsg = true

func SetSwitchTimeMsg(value bool) {
	switchTimeMsg = value
}

//******************************************************************************
// 抛异常函数，避免引用fmt
func Panic(format string, v ...any) {
	panic(fmt.Sprintf(format, v...))
}

//******************************************************************************
func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
}

func GetLoaclIp(isV4 bool) string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if isV4 {
				if ipnet.IP.To4() != nil {
					return ipnet.IP.String()
				}
			} else {
				if ipnet.IP.To4() == nil {
					return ipnet.IP.String()
				}
			}
		}
	}
	return ""
}
