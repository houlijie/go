package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 50; i++ {
		fmt.Println("write i=", i)
		ch <- i
		// time.Sleep(time.Second)
	}
	close(ch)
}
func read(ch chan int, ch2 chan bool) {
	// for v := range ch {

	// }
	for {
		time.Sleep(time.Second)
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("read v=", v)
	}
	ch2 <- true
	close(ch2)
}

func main() {
	ch1 := make(chan int, 50)
	ch2 := make(chan bool, 1)

	go write(ch1)
	time.Sleep(time.Second)
	// 读写的频率可以不一致
	go read(ch1, ch2) //  只写没有读会发生阻塞fatal error: all goroutines are asleep - deadlock!
	time.Sleep(time.Second * 10)

	for i := 0; i < 2; i++ {
		go read(ch1, ch2)
	}

	for {
		_, ok := <-ch2
		fmt.Println(ok)
		if ok {
			break
		}
	}
}
