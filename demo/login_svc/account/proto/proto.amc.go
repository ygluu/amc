package cproto

import (
	"lib/amc"
)

func init() {
	amc.RegisterProtoMsg((*NewUserReq)(nil))
	amc.RegisterProtoMsg((*NewUserRes)(nil))
	amc.RegisterProtoMsg((*LoginReq)(nil))
	amc.RegisterProtoMsg((*LoginRes)(nil))
}
