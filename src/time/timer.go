package main

import (
	"fmt"
	"time"
)

func main() {
	// 新建计时器
	t1 := time.NewTimer(5 * time.Second)
	fmt.Println(t1)
	fmt.Println(time.Now())
	// ch := t1.C
	// fmt.Println(<-ch)

	// 开启goroutine
	go func() {
		<-t1.C
		fmt.Println("normal end...")
	}()

	// time.Sleep(3 * time.Second)
	flag := t1.Stop()
	if flag {
		fmt.Println("early end..")
	}

	// func After
	/*
		NewTimer相同， 只是返回值不一样
	 */
	ch1 := time.After(3 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(ch1)
	fmt.Println(<-ch1)

}
