package main

import (
	"fmt"
	_ "strconv"
)

var (
	// 全局匿名函数
	fc3 = func(x int, y int) int {
		return x * y
	}
)

// 闭包
func cls1(name string) func(string) string {
	return name
}

// func init() {
// 	// var Num1 = 11
// }

func main() {
	// //
	// fmt.Println("num1=", Num1)
	// var name = ""
	// name = strconv.ParseInt(name)

	// 匿名函数直接用
	var fc1 = func(x int) int {
		return x
	}(1)

	var (
		fc2 = func(y int) int {
			return y + 1
		}
	)
	fmt.Println(fc1)

	fmt.Println(fc2(2))
	fmt.Println(fc3(2, 4))

}
