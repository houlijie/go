package main

import "fmt"

func main() {
	a := "hello world"

	// go字符串类型不能改变
	// a = false //  cannot use false (type bool) as type string in assignment
	a = "bbb" // bbb
	fmt.Println(a)

	// 双引号可以识别转义字符 反引号``原样输出， 可以实现防攻击

	var b = "aa \n bb \n"
	var c = `aa \n bb \n`
	fmt.Println(b, c)

	// 字符串的拼接 +, 当一个拼接的操作很长时，可以分行写， +号必须在上一行
	str1 := "hello"
	str2 := "world"
	d := str1 + str2
	fmt.Println(d) // helloworld
	d += " everyone"
	fmt.Println(d) // helloworld everyone

	d = "hello" + "hello" +
		"hello" + "hello" +
		"hello"
	fmt.Println(d)
}
