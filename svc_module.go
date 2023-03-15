package amc

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

// 响应消息服务的模块

type moduleType = uint

const (
	// 基础模块
	mtBase moduleType = iota
	// 主线程核心模块
	mtCore
	// 主线程普通模块
	mtNormal
	// 单实例线程模块
	mtThreadS
	// 多实例线程模块
	mtThreadM
)

func isThreadModule(typ moduleType) bool {
	return (typ == mtThreadS) || (typ == mtThreadM)
}

type ModuleB struct {
	curActTime int64
	curMsg     IMsg
	curMInfo   *methodInfo
}

func (this *ModuleB) Queue() *RecvQueue {
	return nil
}

func (this *ModuleB) Init() {
}

func (this *ModuleB) Ready() {
}

func (this *ModuleB) Final() {
}

func (this *ModuleB) Type() moduleType {
	return mtBase
}

func (this *ModuleB) TimeInfo() *timeInfo {
	return TimeInfo
}

func (this *ModuleB) CurTime() *time.Time {
	return &CurTime
}

func (this *ModuleB) CurTick() int64 {
	return CurTick
}

// 线程和进程间发送消息，线程也可以给自己发送异步消息
func (this *ModuleB) SendMsg(ctx ICtx, msg IMsg) {
	// 由Sender进行分发处理
	if send_queue != nil {
		send_queue.Push(ctx, msg)
	}
}

// 普通模块，由主线程驱动消息的服务方法
type ModuleN struct {
	ModuleB
}

func (this *ModuleN) Queue() *RecvQueue {
	LogI("dddddddd:%-v", recv_queue)
	return recv_queue
}

func (this *ModuleN) Type() moduleType {
	return mtNormal
}

// 给主线程模块派发同步消息，和给子线程派发异步消息
//   注意：不能进程间派发消息
func (this *ModuleN) DispatchMsg(ctx ICtx, msg IMsg) {
	dispatchMsg(ctx, msg)
}

// 核心模块，由主线程驱动消息的服务方法，优先创建和执行
type ModuleC struct {
	ModuleN
}

func (this *ModuleC) Type() moduleType {
	return mtCore
}

// 单实例线程模块，仅一个go协程驱动消息的服务方法
type mtMsg struct {
	ctx      ICtx
	msgId    uint32
	msg      unsafe.Pointer
	method   method
	methodId uint32
}

type ModuleST struct {
	thread
	ModuleB
	msgQueue *RecvQueue
}

func (this *ModuleST) Queue() *RecvQueue {
	if this.msgQueue == nil {
		this.msgQueue = NewRecvQueue()
	}
	return this.msgQueue
}

func (this *ModuleST) Init() {
	if this.msgQueue == nil {
		this.msgQueue = NewRecvQueue()
	}
}

func (this *ModuleST) Run() {
	procMsgQueue(&this.ModuleB, this.msgQueue)

	if this.msgQueue.GetCount() == 0 {
		time.Sleep(time.Millisecond * time.Duration(idleSleepTime))
	}
}

func (this *ModuleST) Type() moduleType {
	return mtThreadS
}

func GetWatchMsg(this *ModuleB) string {
	if this.curActTime == 0 {
		return ""
	}

	curTick := time.Now().UnixMilli()
	long := curTick - this.curActTime
	if long < watchTime {
		return ""
	}

	return fmt.Sprintf("模块<%s>执行方法<%s>处理消息<%s>耗时过长: %d毫秒",
		this.curMInfo.moduleName,
		this.curMInfo.methodName,
		reflect.TypeOf(this.curMsg).Elem().Name(),
		long)
}

func (this *ModuleST) GetWatchMsg() string {
	return GetWatchMsg(&this.ModuleB)
}

// 多实例线程模块，每一个消息go一个线程驱动消息的服务方法
type ModuleMT struct {
	ModuleST
}

func (this *ModuleMT) Type() moduleType {
	return mtThreadM
}
func (this *ModuleMT) Run() {
	doCall := func(msgNode *RecvNode) {
		callMethod(&this.ModuleB, msgNode.Ctx, msgNode.Msg, msgNode.minfo)
	}

	nodes := this.msgQueue.Pop()
	for {
		node := nodes
		if node == nil {
			break
		}
		nodes = nodes.Next
		go doCall(node)
	}

	if this.msgQueue.GetCount() == 0 {
		time.Sleep(time.Millisecond * time.Duration(idleSleepTime))
	}
}

var disableModuleOfName = make(map[string]bool)

func SetModuleDisable(name string) {
	disableModuleOfName[name] = true
}

type moduleInfo struct {
	newFunc func() IModule
}

var moduleInfoOfName = make(map[string]bool)
var moduleInfos = make([]*moduleInfo, 0)

var modulesOfCore = make([]IModule, 0)
var modulesOfNormal = make([]IModule, 0)

// 注册模块创建函数
func RegisterModule(newFunc func() IModule) {
	info := &moduleInfo{
		newFunc: newFunc,
	}
	moduleInfos = append(moduleInfos, info)
}

// 注册模块模块实例
//    owner：宿主模块，非空是必须是ModuleST或ModuleMT的子类
//    module: 模块实例，如果owner非空则module必须是ModuleB子类（即mtBase类型）
func RegisterModuleIns(owner IModule, module IModule) {
	if module == nil {
		Panic("[amc]svc.RegisterModuleIns => 参数不能为空")
	}

	if owner != nil {
		if !isThreadModule(owner.Type()) {
			Panic("[amc]svc.RegisterModuleIns => 模块<%s-%s>的宿主必须是线程类型模块",
				module.Name(), reflect.TypeOf(module).Elem().String())
		}
		if module.Type() != mtBase {
			Panic("[amc]svc.RegisterModuleIns => 模块<%s-%s>类行必须是mtBase类型",
				module.Name(), reflect.TypeOf(module).Elem().String())
		}
	}

	switch module.Type() {
	case mtCore, mtNormal:
		{
			Panic("[amc]svc.RegisterModuleIns => 主线程模块<%s-%s>请用RegisterModule函数注册",
				module.Name(), reflect.TypeOf(module).Elem().String())
		}
	}

	registerModule(owner, module)
}

func registerModule(owner IModule, module IModule) {
	var queue *RecvQueue
	name := module.Name()
	if owner == nil {
		queue = module.Queue()
	} else {
		queue = owner.Queue()
		name = owner.Name() + "_" + name
	}

	mdlType := reflect.TypeOf(module)
	mdlValue := reflect.ValueOf(module)

	if name == "" {
		Panic("[amc]svc.registerModule => 模块对象<%s>缺少模块名", mdlType.Elem().String())
	}

	if queue == nil {
		Panic("[amc]svc.registerModule => 模块<%s-%s>缺少消息队列", name, mdlType.Elem().String())
	}

	typ := module.Type()
	if typ > mtThreadM {
		Panic("[amc]svc.registerModule => 模块<%s-%s>错误：%d", name, mdlType.Elem().String(), typ)
	}

	switch typ {
	case mtCore:
		{
			modulesOfCore = append(modulesOfCore, module)

		}
	case mtThreadS, mtThreadM:
		{
			AddThread(module.(iThread))
			modulesOfNormal = append(modulesOfNormal, module)
		}
	default:
		{
			modulesOfNormal = append(modulesOfNormal, module)
		}
	}

	typeNames := []string{"基础模块", "主线程核心模块", "主线程普通模块", "单实例子线程模块", "多实例子线程模块"}
	LogI("创建%s：%s(%s)", typeNames[typ], name, reflect.TypeOf(module).String()[1:])

	for i := 0; i < mdlType.NumMethod(); i++ {
		registerMethod(module, mdlType.Elem().Name(), queue, mdlType.Method(i), mdlValue.Method(i))
	}
}

func registerModules() {
	for _, info := range moduleInfos {
		module := info.newFunc()
		if disableModuleOfName[module.Name()] {
			LogI("[amc]svc.RegisterModules => %s模块已经禁用", module.Name())
			continue
		}

		var owner IModule

		switch module.Type() {
		case mtCore, mtNormal:
			{
				owner = main_svc
			}
		default:
			{
				owner = nil
			}
		}
		registerModule(owner, module)
	}
}
