package main

import (
	"fmt"
	"time"
)

/*
	临界资源：指并发环境中多个进程/线程/协程共享的资源
	临界资源处理不当，往往会导致数据不一致的问题
 */

func main()  {
	a := 1
	go func() {
		a = 2
		fmt.Println("goroutine...", a)
	}()

	a = 3

	fmt.Println("main...", a)

	time.Sleep(time.Second * 10)
}
