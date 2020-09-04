package main

import (
	"fmt"
	"sync"
)

func hi() {
	fmt.Println("hello")
	wg.Done()
}

var wg sync.WaitGroup

// 程序创建一个主goroutine
func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		// go hi()
		go func(i int) {
			fmt.Println("i=", i)
			wg.Done()
		}(i) // 不传i会出现重复的问题， 因为用的是外面的i
	}

	//主线程结束了, 协程也会退出
	wg.Wait() //等于所有的goroutine完成后再退出
}
