package main

import "fmt"

func in(ch chan int) {
	for i := 1; i <= 20000; i++ {
		// fmt.Println("in i=", i)
		ch <- i
	}
	close(ch)
}

func out(ch1 chan int, ch2 chan int, ch3 chan bool) {
	for {
		num, ok := <-ch1
		// fmt.Println("out i=", num)
		if !ok {
			break
		}
		for i := 2; i < num; i++ {
			if num%i == 0 {
				break
			}
			ch2 <- num
		}
	}
	// close(ch2) // 不能关闭， 因为别的协程还没完成, 交给主线程close
	fmt.Println("协程已完成")
	ch3 <- true

}
func main() {
	ch1 := make(chan int, 2000)
	ch2 := make(chan int, 20000)
	ch3 := make(chan bool, 4)

	// 开启一个协程写
	go in(ch1)
	// 开启4个协程读
	for i := 1; i <= 4; i++ {
		go out(ch1, ch2, ch3)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-ch3
		}

		close(ch2)
	}()

	for {
		_, ok := <-ch2
		if !ok {
			break
		}
		// fmt.Println("v=", v)
	}
	fmt.Println("主线程退出")
}
