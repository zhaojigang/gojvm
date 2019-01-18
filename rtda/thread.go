package rtda

type Thread struct {
	pc    int    // 程序计数器
	stack *Stack // Java虚拟机栈
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024), // 创建线程的时候创建虚拟机栈
	}
}

// 栈帧入栈
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

// 栈帧出栈
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

// 获取当前栈帧
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

// Getter
func (self *Thread) PC() int {
	return self.pc
}

// Setter
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}
func (self *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return newFrame(self, maxLocals, maxStack)
}
