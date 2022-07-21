package main

import (
	"fmt"
	"time"
)

func main() {
	// dl := time.Now().Add(2000 * time.Millisecond)
	// ctx, cancel := context.WithDeadline(context.Background(), dl)
	// defer cancel()
	// select{
	// 	case ctx.Done()
	//
	// }
}

func XBad(done chan bool) {
	time.Sleep(time.Second)
	done <- true
}

func XGood(done chan bool) {
	time.Sleep(time.Second)
	select {
	case done <- true:
	default:
		return
	}
}

func XTimeOut(f func(chan bool)) error {
	done := make(chan bool, 1)
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout....")
	}
}
