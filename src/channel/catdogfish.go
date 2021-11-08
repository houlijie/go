package main

import (
	"fmt"
	"sync"
)

// 轮流交替打印猫狗鱼

func main() {
	var wg = sync.WaitGroup{}
	
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("cat")
		}
	}()
	
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("dog")
		}
	}()
	
	go func() {
		defer wg.Add(1)
		for i := 0; i < 3; i++ {
			fmt.Println("fish")
		}
	}()
	
	wg.Wait()
}