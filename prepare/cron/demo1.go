package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"log"
	"time"
)

func main() {
	var (
		err  error
		expr *cronexpr.Expression
	)

	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}
	now := time.Now()
	// 下次调度的时间
	nextTime := expr.Next(now)
	fmt.Printf("now:%s, \nnextTime:%s \n", now, nextTime)
	time.AfterFunc(nextTime.Sub(now), func() {
		log.Print("调用了")
	})
	time.Sleep(5 * time.Second)
}
