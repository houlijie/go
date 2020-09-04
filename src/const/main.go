package main

import "fmt"

const (
	name string = "hahah"
	num         = 1 // 可以类型推导
	// 表示给num1 赋值0  num2在num1的基础上+1 num3 在num2的基础上+1
	num1 = iota
	num2
	num3, num4 = iota, iota
	// num5 // error  extra expression in const declaratio
	num5 = iota
	num6 = iota
)

/*
	注意事项
	1、常量不能修改
	2、常量初始化时必须给值
	3、常量只能用以下类型 bool\int系列\float系列\string
	4、语法： const xxx [type] = value
*/
func main() {
	fmt.Println(name, num, num1,
		num2, num3, num4, num5, num6) // hahah 1 2 3 4 4 5 6
}
