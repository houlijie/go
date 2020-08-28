package main

import (
	"fmt"
	_ "strconv" // 如果没有使用一个包， 但是不想去掉， 前面加_表示忽略
)

func main() {

	var a = 1
	b := 2

	fmt.Printf("a: %d, b: %d\n", a, b)
	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("a: %d, b: %d\n", a, b)

	//    var i uint = 97

	// var weeks = i / 7
	// var leftDays = i % 7
	// fmt.Printf("剩余%d 个星期零%d天\n", weeks, leftDays)

	// var temp float32
	// fmt.Scanln(&temp)

	// var newtemp float32
	// newtemp = 5.0 / 9 * (temp - 100)
	// fmt.Printf("华氏温度对应的摄氏%v\n", newtemp)

	// // 获取i的地址 用&
	// fmt.Println("i的值：", i)
	// fmt.Println("i的地址：", &i)

	// // 指针变量存的是第一个地址
	// var ptr *uint = &i // ptr 是一个指针变量 ptr的类型是*uint ptr的值是&i， 本身也有地址
	// fmt.Println("ptr的值: ", ptr)
	// fmt.Println("ptr地址：", &ptr)
	// fmt.Println("ptr指向的值： ", *ptr)

	// var num int
	// fmt.Println("num的地址是：", &num)
	// var ptr2 *int = &num
	// *ptr2 = 10
	// fmt.Println("num的值：", num)
}
