package rtda

type Stack struct {
	maxSize uint   // 最多可容纳多少帧
	size    uint   // 当前栈中已经存放的栈帧数量
	_top    *Frame // 栈顶栈帧的指针
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

// 栈帧入栈
func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil {
		frame.lower = self._top // 组建栈帧链表
	}
	self._top = frame // 将当前栈压入栈顶
	self.size++
}

// 栈帧出栈
func (self *Stack) pop() *Frame {
	if self._top == nil { // 栈为空，虚拟机有bug
		panic("jvm stack is empty")
	}
	top := self._top      // 获取栈顶栈帧
	self._top = top.lower // 将当前栈顶栈帧的下一个元素设置为栈顶元素，即弹出栈顶元素
	self.size--
	return top
}

// 获取当前栈帧
func (self *Stack) top() *Frame {
	if self._top == nil { // 栈为空，虚拟机有bug
		panic("jvm stack is empty")
	}
	return self._top
}

func (self *Stack) IsEmpty() bool {
	return self._top == nil
}
