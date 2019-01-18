package control

import (
	"github.com/zhaojigang/gojvm/instructions/base"
	"github.com/zhaojigang/gojvm/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
