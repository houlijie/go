package main

import (
	"fmt"
	"time"
)

func sayhello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}

func test() {
	// recover 解决协程中的panic
	defer func() {
		// 捕获panic
		if err := recover(); err != nil {
			fmt.Println("error", err)
		}
	}()
	var map1 map[int]string
	// map1 = make(map[int]string, 10)
	map1[0] = "ssss"
	fmt.Println("map1", map1)
}
func main() {
	go sayhello()
	go test()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("main hello")
	}
}
