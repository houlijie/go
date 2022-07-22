package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan int)
	ch1 := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			ch1 <- i
		}
		close(ch)
	}()

	strconv.Itoa()

	for i := 0; i < 11; i++ {
		i, ok := <- ch
		if ok {
			fmt.Println(i)
		}
		fmt.Println("done..")
	}
}
