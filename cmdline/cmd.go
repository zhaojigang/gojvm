package cmdline

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	HelpFlag    bool // java -help
	VersionFlag bool // java -version
	// java [-options] <class> [args]
	// java [-options] -jar <jarFile> [args]
	CpOption   string
	XjreOption string   // 指定jre的目录
	Class      string   // 执行主类
	Args       []string // 附加参数
}

func ParseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.HelpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.CpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.CpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	// 如果 parse 成功，向下执行，如果失败，执行 printUsage
	flag.Parse()

	args := flag.Args() // 解析剩余参数
	if len(args) > 0 {
		cmd.Class = args[0]
		cmd.Args = args[1:]
	}
	return cmd
}

func PrintUsage() {
	fmt.Printf("Usage: %s: [-options] class [args...]\n", os.Args[0])
}
