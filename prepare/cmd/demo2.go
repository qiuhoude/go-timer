package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 生成 cmd
	cmd := exec.Command("/bin/bash", "-c", "ls -l")
	// 执行命令,捕获子进程的输出 pipe
	if output, err := cmd.CombinedOutput(); err == nil {
		// 打印子进程输出
		fmt.Printf("%s", output)
	}

}
