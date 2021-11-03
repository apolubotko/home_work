package main

import (
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {

	for env, val := range env {
		_, present := os.LookupEnv(env)
		if present {
			os.Unsetenv(env)
		}
		if val.NeedRemove {
			if err := os.Unsetenv(env); err != nil {
				fmt.Printf("error: %v\n", err)
				return -1
			}
			continue
		}
		// fmt.Printf("Result: %s:%s\n", env, val.Value)
		if err := os.Setenv(env, val.Value); err != nil {
			fmt.Printf("error: %s %v\n", env, err)
			return -1
		}
	}

	app := cmd[0]
	args := cmd[1:]
	command := exec.Command(app, args...)
	out, err := command.CombinedOutput()
	if err != nil {
		fmt.Printf("Executing err: %v\n", err)
		return -1
	}
	exitCode := command.ProcessState.ExitCode()

	fmt.Print(string(out))

	return exitCode
}
