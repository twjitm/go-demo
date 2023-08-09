package multit

import (
	"context"
	"fmt"
	"time"
)

func RunContext(ctx context.Context) {
	go RequestHandler(ctx)
}
func RequestHandler(ctx context.Context) {
	go WaitRedis(ctx)
	go WaitDataBase(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("request done")
		return
	default:
		fmt.Println("run request default")
		time.Sleep(time.Second * 2)
	}
}

func WaitRedis(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("redis done")
			return
		default:
			fmt.Println("run redis default")
			time.Sleep(time.Second * 2)
		}
	}

}
func WaitDataBase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("database done")
			return
		default:
			fmt.Println("run database default")
			time.Sleep(time.Second * 2)
		}
	}
}
