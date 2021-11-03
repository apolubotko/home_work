package main

import (
	"fmt"
	"os"
)

func main() {
	var prog, dir, app string
	var args []string

	// $ go-envdir /path/to/env/dir command arg1 arg2
	for idx, arg := range os.Args {
		switch idx {
		case 0:
			prog = arg
		case 1:
			dir = arg
		case 2:
			app = arg
		default:
			args = append(args, arg)
		}
	}
	checkArgs(prog)
	checkArgs(dir)
	checkArgs(app)
	checkArgs(args)

	envs, err := ReadDir(dir)
	if err != nil {
		fmt.Printf("Something goes wrong: %v\n", err)
		return
	}
	cmd := make([]string, 0)
	cmd = append(cmd, app)
	cmd = append(cmd, args...)
	RunCmd(cmd, envs)
}

func checkArgs(s interface{}) {
	switch s := s.(type) {
	case string:
		if len(s) > 0 {
			return
		}
	case []string:
		if len(s) > 0 {
			return
		}
	default:
		fmt.Printf("wrong arg %#v", s)
		os.Exit(1)
	}

}
