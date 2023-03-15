package sproto

import (
	"lib/amc"
)

func init() {
	amc.RegisterProtoMsg((*UserLoginReq)(nil))
	amc.RegisterProtoMsg((*GetSessionReq)(nil))
	amc.RegisterProtoMsg((*OnLogin)(nil))
}
