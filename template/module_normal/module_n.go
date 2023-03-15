package module

import (
	"lib/amc"
)

// amc.ModuleN：
//	正常模块，服务主线程执行Init、Ready、Uninit和消息方法，和线程模块同级
//  同一等级的模块根据引用关系自动排列执行Init、Ready和倒序执行Uninit

const MName = "普通模块"

type Module struct {
	amc.ModuleN
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
