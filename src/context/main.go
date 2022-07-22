package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	dl := time.Now().Add(2000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), dl)
	defer cancel()
	select{
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(time.Microsecond * 2):
		fmt.Println("timeout...")
	}
}
bps8yRXJiqf6IzI6