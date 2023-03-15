package amc

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"reflect"
	"runtime/debug"
	"unsafe"
)

type eface struct {
	typ  uintptr
	data unsafe.Pointer
}

type method func(ctx ICtx, msg unsafe.Pointer)

type methodInfo struct {
	moduleName string
	methodName string
	msgId      uint32
	method     method
	msgQueue   *RecvQueue
}

var methodInfosOfId = make(map[uint32][]*methodInfo)

func callMethod(module *ModuleB, ctx ICtx, msg IMsg, minfo *methodInfo) {
	defer func() {
		module.curActTime = 0
		err := recover()
		if err != nil {
			LogE("[amc]callMethod => 方法<%s>处理消息<%s>异常：%v\r\n%s", minfo.methodName,
				reflect.TypeOf(msg).Elem().Name(), err, string(debug.Stack()))
			return
		}
	}()

	module.curActTime = CurTick
	module.curMsg = msg
	module.curMInfo = minfo
	minfo.method(ctx, (*eface)(unsafe.Pointer(&msg)).data)
}

func procMsgQueue(module *ModuleB, msgQueue *RecvQueue) {
	nodes := msgQueue.Pop()
	for {
		node := nodes
		if node == nil {
			break
		}
		nodes = nodes.Next
		// 服务方法调用（派发消息）
		callMethod(module, node.Ctx, node.Msg, node.minfo)
	}
}

// 给主线程模块派发同步消息，和给子线程派发异步消息
//   注意：不能进程间派发消息
func dispatchMsg(ctx ICtx, msg IMsg) {
	module := &main_svc.ModuleB
	infos := methodInfosOfId[MsgIdOfMsg(msg)]
	for _, info := range infos {
		queue := info.msgQueue
		if queue == recv_queue {
			// 主线程模块的消息理解调用服务方法，同步派发
			callMethod(module, ctx, msg, info)
		} else {
			// 子线程模块的消息，加入其消息队列中，异步派发
			queue.Push(ctx, msg, info)
		}
	}
}

func isStructPrt(typ reflect.Type) bool {
	return (typ.Kind() == reflect.Ptr) && (typ.Elem().Kind() == reflect.Struct)
}

func registerMethod(module IModule, structName string, msgQueue *RecvQueue, methodStruct reflect.Method, methodValue reflect.Value) {
	methodIface := methodValue.Interface()
	methodType := reflect.TypeOf(methodIface)

	if methodType.Kind() != reflect.Func {
		Panic("[amc]registerMethod => 参数methodValue不是方法函数")
	}

	//methodName := runtime.FuncForPC(methodStruct.Func.Pointer()).Name()
	methodName := structName + "." + methodStruct.Name

	if methodType.NumIn() != 2 {
		return
	}

	if methodType.NumOut() != 0 {
		return
	}

	ctxType := methodType.In(0)

	if ctxType.Name() != "ICtx" {
		return
	}

	msgType := methodType.In(1)
	if !isStructPrt(msgType) {
		return
	}

	if !isMsgObj(msgType) {
		return
	}

	id := MsgIdOfName(msgType.Elem().Name())
	if id == 0 {
		id = registerMsg(reflect.New(msgType).Elem().Interface().(IMsg))
	}

	if isThreadModule(module.Type()) && (id == MsgIdOfMsg(SysMsg_Run)) {
		Panic("[amc]registerMethod => 线程模块<%s>不能实现消息<%s>监听方法",
			reflect.TypeOf(module).Elem().String(), msgType.Elem().Name())
	}

	methodElem := (*eface)(unsafe.Pointer(&methodIface)).data
	method := *(*method)(unsafe.Pointer(&methodElem))

	infos := methodInfosOfId[id]

	if (infos != nil) && (!isBroadcastMsg(msgType)) {
		Panic("[amc]registerMethod => 消息<%s>注册到服务方法<%s>失败，该消息<%s>已经被<%s>注册，非广播类消息不能重复注册服务方法",
			msgType.Elem().Name(), methodName, MsgNameOfId(id), infos[0].methodName)
	}

	info := &methodInfo{
		moduleName: module.Name(),
		methodName: methodName,
		msgId:      id,
		msgQueue:   msgQueue,
		method:     method,
	}

	infos = append(infos, info)
	methodInfosOfId[id] = infos

	return
}
