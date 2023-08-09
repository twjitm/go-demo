package multit

import (
	"context"
	"fmt"
)

func G1(ctx context.Context, name string) {
	fun := func(ctx context.Context, name string) {
		fmt.Println(name)
		G2(ctx, name+"2")
	}
	go fun(ctx, name)

}
func G2(ctx context.Context, name string) {
	go func(ctx context.Context, name string) {
		fmt.Println(name)
	}(ctx, name)
}
