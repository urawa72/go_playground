package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(3 * time.Second))
	defer cancel()

	go sayDeadLine(ctx)

	select {
	case <- ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func sayDeadLine(ctx context.Context) {
	for {
		fmt.Println(ctx.Deadline())
		time.Sleep(1 * time.Second)
	}
}
