package main

import (
	. "byx-script-go/common"
	. "byx-script-go/interpreter"
	"flag"
	"fmt"
)

type ImportPaths []string

func (p ImportPaths) String() string {
	return "./"
}

func (p *ImportPaths) Set(path string) error {
	*p = append(*p, path)
	return nil
}

// ByxScript命令行工具
func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	paths := ImportPaths{"./"}
	flag.Var(&paths, "importPath", "import paths")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("please specify the script file to run")
		return
	}

	scriptPath := args[0]
	script, err := ReadFileAsString(scriptPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	RunScript(script, RunConfig{ImportPaths: paths})
}
