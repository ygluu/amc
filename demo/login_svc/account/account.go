package account

import (
	"lib/amc"
	"lib/amc/demo/login_svc/account/proto"
	"lib/amc/demo/proto"
)

// amc.ModuleN：
//	线程模块，服务主线程执行Init、Ready、Uninit，消息由服务主线程传递，子线程执行消息方法，和普通模块同级
//  同一等级的模块根据引用关系自动排列执行Init、Ready和倒序执行Uninit
const MName = "账号管理"

type Module struct {
	amc.ModuleMT // 多线程并非执行消息方法，即每个消息方法都go一次
}

func init() {
	amc.RegisterModule(func() amc.IModule { return &Module{} })
}

func (m *Module) Name() string {
	return "账号管理"
}

func (m *Module) OnSM_SecChg(ctx amc.ICtx, msg *amc.SysMsgSecChg) {

}

func (m *Module) OnCM_Login(ctx amc.ICtx, req *sproto.UserLoginReq) {
	// 验证登录，基于amc.ModuleMT多并发线程类，
	// 在这里可以直接读写数据库或者进行Http请求
	// ...

	// 广播给所有监听OnLogin的服务，OnLogin定义为广播消息
	onLogin := &sproto.OnLogin{}

	onLogin.Name = req.Name
	onLogin.Token = amc.GetGUID()
	onLogin.NameHash = amc.LB.GetHash(onLogin.Name)
	_, onLogin.OwnerAddr = amc.LB.GetInfoByHash(amc.MsgIdOfMsg(onLogin), onLogin.NameHash)
	onLogin.GateAddr = req.GateAddr
	onLogin.GateClient = req.GateClient

	m.SendMsg(amc.NullCtx, onLogin)

	// Write session info to redis
	// redis.Write(onLogin.Token, onLogin)

	// 告诉客户端登录结果
	res := &cproto.LoginRes{
		Ret:   0,
		Msg:   "登录成功",
		Token: onLogin.Token,
	}
	m.SendMsg(ctx, res)
	amc.LogI("OnCM_Login => Req:%s, Res:%s", req, res)
}

func (m *Module) OnCM_NewUser(ctx amc.ICtx, req *cproto.NewUserReq) {
	// 告诉客户端登录结果
	res := &cproto.NewUserRes{
		Ret: 0,
		Msg: "注册成功",
	}
	m.SendMsg(ctx, res)
	amc.LogI("OnCM_NewUser => Req:%s, Res:%s", req, res)
}

func (m *Module) Init() {

}

func (m *Module) Uninit() {

}

func (m *Module) Ready() {

}
