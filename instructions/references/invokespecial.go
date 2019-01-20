package references

import (
	"github.com/zhaojigang/gojvm/instructions/base"
	"github.com/zhaojigang/gojvm/rtda"
)

type INVOKE_SPECIAL struct{ base.Index16Instruction }

// hack
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}
