package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 生成cmd
	cmd := exec.Command("/bin/bash", "-c", "open ./")

	err := cmd.Run()
	fmt.Print(err)
}
