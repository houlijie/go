package main

import (
	"fmt"
	"time"
)

/*
	布隆过滤器c
*/

func main()  {
	chan2 := make(chan int)
	go func() {
		<-chan2
	}()
	close(chan2)
	select {
	case w := <-chan2:
		 fmt.Println("world ", w)
	}
	time.Sleep(time.Second * 33)
}

