package interpreter

import (
	"fmt"
	"github.com/zhaojigang/gojvm/instructions"
	"github.com/zhaojigang/gojvm/instructions/base"
	"github.com/zhaojigang/gojvm/rtda"
	"github.com/zhaojigang/gojvm/rtda/heap"
)

func Interpret(method *heap.Method) {
	thread := rtda.NewThread()                    // 创建线程(同时创建线程私有的stack)
	frame := thread.NewFrame(method) // 创建栈帧(同时创建栈帧私有的localVariableTable和operandStack)
	thread.PushFrame(frame)                       // 将栈帧push到线程的stack中

	// 每个方法的执行结束必须是return指令
	// 此处先catch错误
	defer catchErr(frame)
	loop(thread, method.Code())
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()       // 弹出顶部frame
	reader := &base.ByteCodeReader{} // 创建字节码读取器

	for {
		pc := frame.NextPC() // 获取 frame 中将要读取的下一字节的位置 nextPc
		thread.SetPC(pc)     // 将该 nextPc 设置给线程 pc

		// decode
		reader.Reset(bytecode, pc)
		opCode := reader.ReadUint8()                       // 读取操作码 opCode（指令类型）
		instruction := instructions.NewInstruction(opCode) // 根据opCode创建相应的指令
		instruction.FetchOperands(reader)                  // 从字节码中读取操作数
		frame.SetNextPC(reader.PC())                       // 将当前读取到的字节码的位置设置到 frame 的 nextPc 中，用于执行下一条指令

		// execute
		fmt.Printf("pc:%2d, inst:%T %v\n", pc, instruction, instruction)
		instruction.Execute(frame) // 执行字节码指令
	}
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}
