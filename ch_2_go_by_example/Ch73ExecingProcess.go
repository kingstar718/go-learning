package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}
	// exec还需要一组环境变量，这里只提供我们当前的环境
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
