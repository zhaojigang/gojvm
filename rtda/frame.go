package rtda

// 栈帧
type Frame struct {
	lower        *Frame        // 下一个栈帧
	localVars    LocalVars     // 本地变量表
	operandStack *OperandStack // 操作数栈
}

// 执行方法所需的局部变量大小 maxLocals 和操作数栈深度会由编译器预先计算好，
// 存储在 class 文件的 method_info 方法表结构中的 code 属性中
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
