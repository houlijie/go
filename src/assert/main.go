package main

import "fmt"

type h struct {
	name string
}

// 类型断言
/*
	场景： 由于接口是一般类型， 不知道具体类型， 如果要专成具体类型 就需要使用类型断言
*/
func main() {
	var a1 interface{}
	var b1 h = h{name: "hh"}

	a1 = b1 // 空接口，可以接收任何类型
	fmt.Println(a1)

	// b1 = a1 // error: cannot use a1 (type int) as type string in assignment
	b1 = a1.(h)     // 这里就是类型断言, 这里能成功的原因是  a1本来就是h类型
	fmt.Println(b1) // {hh}

	var a2 interface{}
	var b2 float64 = 4
	// var b3 int = 10
	a2 = b2

	b3, ok := a2.(float32)
	if ok {
		fmt.Println("断言成功")
	} else {
		fmt.Println("断言失败")
	}
	fmt.Println(b3)

}
