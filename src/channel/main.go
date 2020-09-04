package main

import "fmt"

/*
	channel 引用
	1、本质是队列， 先进先出
	2、线程安全，多goroutine访问时 不需要加锁
	3、channel有类型，只能存放定义类型的数据
	4、管道可以声明为只读或只写
	5、使用select可以解决从管道取数据的阻塞问题 eg: exec05
	6、使用recover解决协程中panic eg: exec06
	7、无缓冲的通道，只有有人接收时才能发送

	关闭channel
	1、继续往已关闭的通道再发送值会导致panic
	2、关闭一个已经关闭的通道会导致panic
	3、对一个已经关闭并且没值的通道执行接收操作会得到对应类型的零值
	4、关闭通道不是必须的，它可以被垃圾机制回收

*/
// 只读
func send(ch <-chan int) {

}

// 只写
func wr(ch chan<- int) {

}

func main() {
	var ch chan int                                //双向管道
	ch = make(chan int, 4)                         // 无缓冲通道，必须有接收才能发送
	fmt.Printf("ch = %v, ch addr = %v\n", ch, &ch) // 地址 0xc00006c000  0xc000122018

	// 只读
	// var chreadonly chan<- int
	// chreadonly = make(chan int, 3)
	// chreadonly <- 10 // error

	// 只写
	// var chwriteonly <-chan int
	// chwriteonly = make(chan int, 3)
	// num := <-chwriteonly // error
	// fmt.Println(num)

	// 写入
	ch <- 10
	numtes := 22
	ch <- numtes
	ch <- 11
	ch <- 23
	// 增加的数量不能超过cap
	// ch <- 33   // fatal error: all goroutines are asleep - deadlock!
	fmt.Printf("len=%d, cap=%d\n", len(ch), cap(ch)) // len=2, cap=3

	// 取数据
	num1 := <-ch // 10
	num2 := <-ch // 22
	num3 := <-ch // 11
	fmt.Printf("num1=%d, num2=%d, num3=%d\n", num1, num2, num3)

	// 在没有使用协程的情况下 数据已取出再继续取， 会报错
	// num4 := <-ch
	// fmt.Println("num4=", num4) // fatal error: all goroutines are asleep - deadlock!

	// 关闭，关闭后不能继续写入 可以继续读取
	close(ch)
	// ch <- 10 // panic: send on closed channel
	num4 := <-ch
	fmt.Println("num4=", num4) // 23 关闭后还可以继续读取

	// channel遍历 for range,不能使用for, 因为len每次取出都会变化，取出类似pop
	/*
		注意事项
		1、遍历前管道必须先close
		2、
	*/
	ch2 := make(chan int, 3)
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	close(ch2)           // fatal error: all goroutines are asleep - deadlock!
	close(ch2)           // panic: close of closed channel
	for v := range ch2 { // notice: 管道没有index
		fmt.Println(v)
	}
}
