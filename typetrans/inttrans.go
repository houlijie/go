package main

import (
	"fmt"
)

// 整型类型转换
func main() {
	// 类型转换
	var a int32 = 6999
	fmt.Println(a)
	// 注意事项
	// 1、被转换的事变量存储的数据， 对源数据本身的数据类型没有变化
	var b int64 = int64(a)
	fmt.Println(b) // 6999

	// 2、范围大-->范围小转换过程中， 编译不会报错， 但是结果按照溢出处理
	var c int8 = int8(a)
	fmt.Println(c) // 87

	var n1 int32 = 12
	var n2 int8
	// var n3 int64

	n2 = int8(n1) + 10 //不同类型变量赋值时，必须要先进行类型转换
	fmt.Println(n2)

	n3 := n1 + 20 // 编译不通过，cannot use n1 + 20 (type int32) as type int64 in assignment 因为 n1 + 20 为int32类型，n3为int64类型
	fmt.Println(n3)

	var n4 int8
	n4 = int8(n1) + 127
	fmt.Println(n4) // -127 能编译通过，但是溢出

	n5 := int8(n1) + 128 // constant 128 overflows int8 编译不通过
	fmt.Println(n5)
}
