package amc

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

type logType = int

const (
	lt_E logType = iota
	lt_W
	lt_I
	lt_D
)

var tnames = []string{"E", "W", "I", "D"}

func doLog(typ logType, format string, v ...interface{}) {
	logger("| "+tnames[typ]+" | "+cluster_name+" | "+svc_name+" | "+svc_addr+" | "+format+"\n", v...)
}

func LogE(format string, v ...interface{}) {
	doLog(lt_E, format, v...)
}

func LogW(format string, v ...interface{}) {
	doLog(lt_W, format, v...)
}

func LogI(format string, v ...interface{}) {
	doLog(lt_I, format, v...)
}

func LogD(format string, v ...interface{}) {
	if isDebugMode {
		doLog(lt_D, format, v...)
	}
}
