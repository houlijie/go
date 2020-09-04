package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
	协程特点：
	1、具有独立的栈空间
	2、共享程序堆空间
	3、调度由用户控制
	4、协程是轻量级的线程

	结论：
	1、主线程是一个物理线程， 直接作用在cpu上， 是重量级的，非常消耗cpu资源
	2、协程是主线程开启的， 是轻量级的线程 ， 是逻辑态，对资源消耗相对较小

	MPG模式
	1、M 主线程
	2、P 协程执行需要的上下文
	3、G 协程

	不同goroutine如何通讯
	1、全局变量的互斥锁
	2、使用管道channel解决
*/
func test(max int) {
	for i := 0; i < max; i++ {
		fmt.Println("test i=", i)
	}
	time.Sleep(time.Second)
}

func test02() {
	time.Sleep(time.Second)
	fmt.Println("test02")
}
func main() {
	cpunum := runtime.NumCPU()
	fmt.Println("cpunum=", cpunum)
	go test02()
	go test(10) // 主线程退出，协程也会退出(不管是否完成)

	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println("main i=", i)
	}

}
