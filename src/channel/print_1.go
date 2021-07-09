package main

import (
	"fmt"
	"sync"
	_ "sync"
	_ "time"
)

func odd(x int, ch chan int)  {
	defer wg.Done()
	for i := 1; i <= x; i++ {
		ch<-1
		if i % 2 == 1 {
			fmt.Println("i = ", i)
		}
	}
}

func even(x int, ch chan int)  {
	defer wg.Done()
	for i := 1; i <= x; i++ {
		<-ch
		if i % 2 == 0 {
			fmt.Println("i = ", i)
		}
	}
}

var wg sync.WaitGroup

func main()  {
	ch := make(chan int)

	wg.Add(2)

	x := 10
	go odd(x, ch)
	go even(x, ch)

	wg.Wait()
}
