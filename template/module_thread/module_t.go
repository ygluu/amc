package module

import (
	"lib/amc"
)

// amc.ModuleN：
//	线程模块，服务主线程执行Init、Ready、Uninit，消息由服务主线程传递，子线程执行消息方法，和普通模块同级
//  同一等级的模块根据引用关系自动排列执行Init、Ready和倒序执行Uninit
const MName = "线程模块"

// 注意：本模块的消息方法是子线程执行的
type Module struct {
	amc.ModuleST // 单线程并非执行消息方法
	//amc.ModuleMT // 多线程并非执行消息方法，即每个消息方法都go一次
}

func init() {
	amc.RegisterModule(MName, func() amc.IModule { return &Module{} })
}

func (m *Module) Init() {

}

func (m *Module) Uninit() {

}

func (m *Module) Ready() {

}

func (m *Module) OnSM_Ready(ctx amc.ICtx, msg *amc.SysMsgSecChg) {

}
