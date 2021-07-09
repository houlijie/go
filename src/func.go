package main

import (
	"fmt"
	"strconv"
	_ "strconv"
)

// 有返回的函数
func fc1() int {
	x := 1
	return x
}

// 命名返回值，相当于在函数声明一个变量
func fc2() (res int) {
	x := 1
	res += x

	return // 等同于return res, 因为上面已经定义了(res int)
}

// 多个返回值
func fc3() (int, string) {
	x := 1
	y := "hahaah"
	return x, y
}

// 参数的类型简写
func fc4(x, y int, m, n string) (int, string) {
	z := "hahaah"
	return x + y, z
}

// 可变参数
func f5(y ...int) { // 可变参数必须写在参数的最后
	fmt.Println(y) // y是slice类型 []int和函数上的定义保持一致
}

// 没有默认参数概念

var (
	// 全局匿名函数
	fc6 = func(x int, y int) int {
		return x * y
	}
)

// 闭包
func cls1(name string) func(string) string {
	return name
}

func init() {
	// var Num1 = 11
}

func main() {
	//
	fmt.Println("num1=", Num1)
	var name = ""
	name = strconv.ParseInt(name)

	匿名函数直接用
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
