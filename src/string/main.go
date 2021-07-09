package main

import (
	"fmt"
	_ "strconv"
)

func f1() int {
	x := 5
	defer func() {
		fmt.Println("xxx", x)
		x++
		fmt.Println("xxx", x)
	}()
	return x
}

func f2() (x int) {
	defer func() {
		fmt.Println("xxx", x)
		x++
		fmt.Println("yyy", x)
	}()
	return 7
}

func f3() (y int) {
	x := 5
	defer func() {
		fmt.Println("yyy", x)
		x++
		fmt.Println("yyy", x)
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func main() {
	s := "sss"
	s1 := []rune(s)
	for i, v := range s1 {
		fmt.Printf("%v => %c", i, v)
	}
	// fmt.Println(f4())
	// 常用函数
	// len返回字节数
	// var name = "zhangsan"
	// fmt.Println(len(name)) // 8

	// var name1 = "zhangsan张三"
	// fmt.Println(len(name1)) // 14 返回字节数， 1个汉字3个字节

	// var name2 = []rune(name1)
	// for i := 0; i < len(name2); i++ {
	// 	fmt.Printf("c=%c\t", name2[i])
	// }

	// num1, err := strconv.Atoi(name)
	// if err != nil {
	// 	fmt.Println(num1)
	// } else {
	// 	return
	// }
}
