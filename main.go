package main

import (
	"fmt"
	"github.com/zhaojigang/gojvm/classfile"
	"github.com/zhaojigang/gojvm/classpath"
	"github.com/zhaojigang/gojvm/cmdline"
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
	// 1. 构造 classpath 实例（构造三个Entry）
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	// 2. 寻找并读取要查询的类
	className := strings.Replace(cmd.Class, ".", "/", -1)

	cf := loadClass(className, cp)
	printClassInfo(cf)
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
	fmt.Printf("magic: %v\n", cf.Magic())
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("access flags: %v\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())

	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("   %v\n", f.Name())
	}

	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("   %v\n", m.Name())
	}
}