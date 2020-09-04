package main

import "fmt"

func main() {
	// 使用select可以解决从管道取数据的阻塞问题

	intch := make(chan int, 5)
	charch := make(chan string, 3)
	for i := 0; i < 5; i++ {
		intch <- i
	}

	for i := 0; i < 3; i++ {
		charch <- fmt.Sprintf("hello %d", i)
	}
lb1: // label
	for {
		select {
		// 如果intchan一直没有关闭， 不会一致阻塞而deadlock, 会自动到下一个case匹配
		case v := <-intch:
			fmt.Println("v=", v)
		case v := <-charch:
			fmt.Println("char=", v)
		default:
			fmt.Println("结束")
			// return // return main
			break lb1
		}
	}

	fmt.Println("main 结束") // break lb1只是结束循环
}
