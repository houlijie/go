package main

import "fmt"

func chtest1() {
	ch := make(chan int)
	ch <- 10
	fmt.Println("发送成功")
}

func chgoroutine(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func chtest2() {
	ch := make(chan int)
	go chgoroutine(ch) // 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}
func main() {
	chtest1()
	chtest2()
}
