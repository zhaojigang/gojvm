package stores

import (
	"github.com/zhaojigang/gojvm/instructions/base"
	"github.com/zhaojigang/gojvm/rtda"
)

type ASTORE struct{ base.Index8Instruction }
type ASTORE_0 struct{ base.NoOperandsInstruction }
type ASTORE_1 struct{ base.NoOperandsInstruction }
type ASTORE_2 struct{ base.NoOperandsInstruction }
type ASTORE_3 struct{ base.NoOperandsInstruction }

func _astore(frame *rtda.Frame, index uint) {
	ref := frame.OperandStack().PopRef() // 从操作数栈pop
	frame.LocalVars().SetRef(index, ref) // 放到局部变量表中指定的索引index位置处

}

func (self *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, self.Index)
}

func (self ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}

func (self *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

func (self *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

func (self *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}
