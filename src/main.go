package main

import (
	"context"
	"fmt"
	"go-demo/src/logic/multit"
	"time"
)

func main() {
	//start server
	ctx := context.Background()

	multit.G1(ctx, "wenjiang.tang")
	ctx.Deadline()
	time.Sleep(time.Second)

	ctx, cancel := context.WithCancel(ctx)
	//ctx, cancel := context.WithDeadline(ctx, time.Now().Add(6*time.Second))
	//context.WithTimeout()
	fmt.Println("ran before")
	multit.RunContext(ctx)
	fmt.Println("ran after")
	time.Sleep(5 * time.Second)

	cancel() //发送关闭指令

	fmt.Println("cancel ****")
	time.Sleep(5 * time.Second)
}
