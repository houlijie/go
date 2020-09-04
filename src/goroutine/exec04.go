package main

import (
	"fmt"
	"runtime"
)

func g1() {
	defer fmt.Println("A.defer")
	func() {
		defer fmt.Println("B.defer")
		// 结束协程
		defer fmt.Println("C.defer")
		fmt.Println("B")
	}()
	runtime.Goexit() //终止调用它的go程
	fmt.Println("A")
}
func g2() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

/*
	Goexit终止调用它的go程。其它go程不会受影响。Goexit会在终止该go程前执行所有defer的函数。
*/
func main() {
	go g2()
	go g1()
	for {
	}
}
