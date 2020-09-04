package main

import (
	"context"
	"time"
)

func main() {

	dl := time.Now().Add(2000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), dl)
	defer cancel()
	select{
		case ctx.Done()
		
	}
}
