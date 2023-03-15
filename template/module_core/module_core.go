package module

import (
	"lib/amc"
)

// amc.ModuleC：
//	核心模块优于其他模块执行：Init、Ready
//	核心模块次于其他模块执行：Uninit
//  同一等级的模块根据引用关系自动排列执行Init、Ready和倒序执行Uninit

const MName = "核心模块"

type Module struct {
	amc.ModuleC
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
