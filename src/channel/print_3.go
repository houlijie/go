package main

import (
	"fmt"
	"runtime"
	"sync"
	_ "time"
)

func ping(x int,  out chan int)  {
	defer wg1.Done()
	for {
		v := <-out
		fmt.Println("ping > ", v)
		runtime.Gosched()
		out <- v+1
	}
}

func pong(x int, out chan int)  {
	defer wg1.Done()
	for {
		v := <-out
		fmt.Println("pong > ", v)
		runtime.Gosched()
		out <- v+1
	}
}

var wg1 sync.WaitGroup

func main()  {
	ch := make(chan int)

	wg1.Add(2)
	x := 10
	go ping(x, ch)
	go pong(x, ch)
	ch <- 1

	wg1.Wait()
}
