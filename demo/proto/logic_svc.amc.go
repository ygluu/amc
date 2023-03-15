package sproto

import (
	"lib/amc"
)

func init() {
	amc.RegisterProtoMsg((*Login)(nil))
}
