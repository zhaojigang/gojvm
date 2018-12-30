package main

import (
	"fmt"
	"github.com/zhaojigang/gojvm/classfile"
	"github.com/zhaojigang/gojvm/classpath"
	"github.com/zhaojigang/gojvm/cmdline"
	"github.com/zhaojigang/gojvm/rtda"
	"strings"
)

func main() {
	cmd := cmdline.ParseCmd()
	if cmd.VersionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.HelpFlag || cmd.Class == "" {
		cmdline.PrintUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *cmdline.Cmd) {
	//1. 构造 classpath 实例（构造三个Entry）
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	// 2. 寻找并读取要查询的类
	className := strings.Replace(cmd.Class, ".", "/", -1)

	cf := loadClass(className, cp)
	printClassInfo(cf)

	//frame := rtda.NewFrame(100, 100)

	//testLocalVars(frame.LocalVars())
	//testOperandStack(frame.OperandStack())
}

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)

	println(vars.GetInt(0))    // 100
	println(vars.GetInt(1))    // -100
	println(vars.GetLong(2))   // 2997924580
	println(vars.GetLong(4))   // -2997924580
	println(vars.GetFloat(6))  // +3.141593e+000
	println(vars.GetDouble(7)) // +2.718282e+000
	println(vars.GetRef(9))    // 0x0
}

func testOperandStack(stack *rtda.OperandStack) {
	stack.PushInt(100)
	stack.PushInt(-100)
	stack.PushLong(2997924580)
	stack.PushLong(-2997924580)
	stack.PushFloat(3.1415926)
	stack.PushDouble(2.71828182845)
	stack.PushRef(nil)

	println(stack.PopRef())    // 0x0
	println(stack.PopDouble()) // +2.718282e+000
	println(stack.PopFloat())  // +3.141593e+000
	println(stack.PopLong())   // -2997924580
	println(stack.PopLong())   // 2997924580
	println(stack.PopInt())    // -100
	println(stack.PopInt())    // 100
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	// 寻找类文件，并直接读取为[]byte
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	// 将 []byte 按照 class 格式解析成 ClassFile 对象
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("magic: %v\n", cf.Magic())                                // magic: 0
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion()) //version: 52.0
	fmt.Printf("access flags: %v\n", cf.AccessFlags())                   //access flags: 49
	fmt.Printf("this class: %v\n", cf.ClassName())                       //this class: java/lang/String
	fmt.Printf("super class: %v\n", cf.SuperClassName())                 //super class: java/lang/Object
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())                  //interfaces: [java/io/Serializable java/lang/Comparable java/lang/CharSequence]

	fmt.Printf("fields count: %v\n", len(cf.Fields())) // fields count: 5
	for _, f := range cf.Fields() {
		fmt.Printf("   %v\n", f.Name())
	}

	fmt.Printf("methods count: %v\n", len(cf.Methods())) //methods count: 94
	for _, m := range cf.Methods() {
		fmt.Printf("   %v\n", m.Name())
	}
}
