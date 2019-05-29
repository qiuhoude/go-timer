package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type reslut struct {
	err    error
	output []byte
}

func main() {
	// 执行一个cmd ,让它在一个协程里去执行,让它执行2秒 sleep 2; echo hi, 1秒的时候,用代码kill掉
	var ()

	// 创建一个结果队列
	reslutChan := make(chan *reslut)

	// ctx原理   表示 ctx 继承与context.TODO_()
	// context 里面有个 chan struct{}
	// cancelFunc 调用的时候,就是 close(chan)
	ctx, cancelFunc := context.WithCancel(context.TODO())

	go func() {
		cmd := exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 2,echo hi")
		// 执行命令，等待结果返回
		output, err := cmd.CombinedOutput()
		// 通过chan输出结果
		reslutChan <- &reslut{
			err:    err,
			output: output,
		}
	}()

	time.Sleep(1 * time.Second)
	// 取消上下文
	cancelFunc()
	// 在main协程中 等待子协程的退出
	res := <-reslutChan
	fmt.Printf("%s, %s\n", res.err, res.output)
}
