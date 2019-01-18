package interpreter

import (
	"fmt"
	"github.com/zhaojigang/gojvm/classfile"
	"github.com/zhaojigang/gojvm/instructions"
	"github.com/zhaojigang/gojvm/instructions/base"
	"github.com/zhaojigang/gojvm/rtda"
)

func Interpreter(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals() // 获取局部变量表所需的空间大小
	maxStack := codeAttr.MaxStack()   // 获取操作数栈大小
	bytecode := codeAttr.Code()

	thread := rtda.NewThread()                    // 创建线程
	frame := thread.NewFrame(maxLocals, maxStack) // 创建栈帧
	thread.PushFrame(frame)                       // 将栈帧push到线程的JVM虚拟机栈中

	// 每个方法的执行结束必须是return指令
	// 此处先catch错误
	defer catchErr(frame)
	loop(thread, bytecode)
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		instruction := instructions.NewInstruction(opcode)
		instruction.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		// execute
		fmt.Printf("pc:%2d, inst:%T %v\n", pc, instruction, instruction)
		instruction.Execute(frame)
	}
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}
