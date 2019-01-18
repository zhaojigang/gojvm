package base

import "github.com/zhaojigang/gojvm/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader) // 从字节码流中取出操作数
	Execute(frame *rtda.Frame)            // 执行指令逻辑
}

// 1. 没有操作数的指令
type NoOperandsInstruction struct{}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

// 2. 跳转指令
type BranchInstruction struct {
	Offset int // 跳转偏移量
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

// 3. 存储和加载类指令需要根据索引存取局部变量表，索引由单字节(8位)操作数给出
type Index8Instruction struct {
	Index uint // 局部变量表索引
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

// 4. 有一些指令需要访问运行时常量池，常量池索引由2个字节(8位)操作数给出
type Index16Instruction struct {
	Index uint // 常量池索引
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
