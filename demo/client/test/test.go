package test

import (
	"lib/amc"
	"lib/amc/demo/login_svc/account/proto"
	"lib/amc/plugin/lb"
)

var mynet amc.INet
var mylb *lb.HashBalan
var mycdc amc.ICodec

func Start(net amc.INet, lb *lb.HashBalan, codec amc.ICodec) {
	mynet = net
	mylb = lb
	mycdc = codec
}

type test struct {
	amc.ModuleN
	isSendLogin bool
}

func (this *test) Name() string {
	return "压力测试模块"
}

func (this *test) OnMsg_Login(ctx amc.ICtx, msg *cproto.LoginRes) {
	amc.LogI("登录成功：%v", msg)
}

func (this *test) OnMsg_SecChg(ctx amc.ICtx, msg *amc.SysMsgSecChg) {
	if this.isSendLogin {
		return
	}

	_, addr := mylb.GetInfo()
	if addr == "" {
		amc.LogI("获取网关地址失败")
		return
	}

	conn, err := mynet.Dial(addr)
	if err != nil {
		amc.LogI("链接目标<%s>失败：%v", addr, err)
		return
	}

	req := &cproto.LoginReq{
		Name:     "abc",
		Password: "123456",
	}
	data, err := mycdc.Marshal(req)
	if err != nil {
		amc.LogI("序列化消息失败：%v", err)
		return
	}

	err = mynet.SendPack(conn, amc.MsgIdOfMsg(req), 0, data)
	if err != nil {
		amc.LogI("发送消息失败：%v", err)
		return
	}

	this.isSendLogin = true
}

func init() {
	amc.RegisterModule(func() amc.IModule { return &test{} })
}
