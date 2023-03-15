package sproto

import (
	"lib/amc"
)

func init() {
	amc.RegisterProtoMsg((*Broadcast)(nil))
}
