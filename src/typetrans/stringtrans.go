package main

import (
	"fmt"
	"strconv"
)

func main() {
	// =========基本类型转字符串========
	var n1 int8 = 1
	var n2 float64 = 1.2
	// var c1 byte = 'h'
	// var b1 bool = true
	var str string

	// 1.fmt.Sprintf("%参数", 表达式)
	// str = fmt.Sprintf("%d", n1)
	// fmt.Printf("str type is %T str=%q \n", str, str) // str type is string str=1

	// str = fmt.Sprintf("%f", n2)
	// fmt.Printf("str type is %T str=%q\n", str, str) //str type is string str=1.200000

	// str = fmt.Sprintf("%c", c1)
	// fmt.Printf("str type is %T str=%q \n ", str, str) //str type is string str=1.200000

	// str = fmt.Sprintf("%t", b1)
	// fmt.Printf("str type is %T str=%q \n ", str, str) //str type is string str=1.200000

	// 2.strconv包函数
	str = strconv.FormatInt(int64(n1), 10)
	fmt.Printf(str) // 1

	str1 := strconv.FormatFloat(n2, 'f', 10, 64)
	fmt.Printf(str1, "\n") // 1

	var num int = 4567
	str = strconv.Itoa(num) // inttostring
	fmt.Printf(str)

	// =========string转基本类型 strconv包=====
	// ParseInt ParseUint ParseFloat ParseBool
	var str2 = "1234"
	var num1 int
	num1 = strconv.ParseInt(str2)
	fmt.Println(num1)
}
