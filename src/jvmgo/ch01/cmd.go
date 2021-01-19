package main

import (
	"fmt"
	"flag"
	"os"
)

type Cmd struct {
	helpFlag     bool
	versionFlag  bool
	cpOption     string
	class        string
	args         []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.BoolVar(&cmd.helpFlag, "help", false, "help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "version message")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath message")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath message")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

// 命令错误时打印该信息
func printUsage() {
	// os.Args 存放了传递给命令行的全部参数
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

