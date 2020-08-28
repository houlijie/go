package main

import (
	"fmt"
)

func sum(x int, y int) int {
	defer fmt.Println("x=", x) // 1
	defer fmt.Println("y=", y) // 2

	total := x + y // 3

	defer fmt.Println("total=", total) // 4

	return total // 5
}

// name := "张三" // notice： 语法错误， 因为这里等同于定义和赋值， 赋值语句不能在函数体外使用

func main() {

	// 场景1
	// 下面执行结果是： total= 7 y= 4 x= 3 num= 7 main test
	// 原因是：
	// defer 延迟执行， 后进先出
	// defer 压入栈时，也会将相应的值copy同时入栈
	// defer fmt.Println("main test") // 6

	num := sum(3, 4) // 7

	fmt.Println("num=", num) // 8

	// 场景2
	// 作为闭包引用的话，则会在defer函数真正调用时根据整个上下文确定当前的值。
	var whatever [2]struct{}

	for i, _ := range whatever {
		defer func() {
			fmt.Println(i) // 执行结果 2 2
		}()
	}
}
