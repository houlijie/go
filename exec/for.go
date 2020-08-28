package main

import (
	"fmt"
)

func main() {
	test1(2)

	// var nNum = 0
	// var total = 0
	// for i := 0; i <= 1 00; i++ {
	// 	if i%9 == 0 {
	// 		nNum++
	// 		total += i
	// 	}
	// }
	// fmt.Printf("9的倍数的整数有%d个，总和是%d \n", nNum, total)
	// b := 6
	// for a := 0; a < 7; a++ {
	// 	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	// 	b--
	// }
	// 打印乘法表
	// for c := 1; c < 10; c++ {
	// 	for d := 1; d < c+1; d++ {
	// 		fmt.Printf("%d * %d = %d \t", d, c, c*d)
	// 	}
	// 	fmt.Println()
	// }

	// 判断是不是水仙花
	// var num string
	// fmt.Printf("请输入一个正整数")
	// fmt.Scanln(&num)
	// var total uint64
	// for j := 0; j < len(num); j++ {
	// 	total += j * j * j
	// 	if total == uint64(num) {
	// 		fmt.Printf("%c是水仙花数")
	// 	}
	// }

	// 指定标签
	// label1: // 设置一个标签
	// 	for h := 0; h < 10; h++ {
	// 	label2:
	// 		for m := 0; m < 2; m++ {
	// 			if m == 1 {
	// 				break label2 // 默认跳出最近的for循环, 指定标签，跳出指定标签的循环
	// 			}
	// 			fmt.Println("m=", m)
	// 			break label1
	// 		}
	// 	}
	//

	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Printf(" 当前的数字是%d, 总数是%d\n ", i, sum)
			break
		}
	}
	嘎嘎啊

	name := ""
	pwd := ""
	for i := 3; i > 0; i-- {
		fmt.Println("请输入账号")
		fmt.Scanln(&name)
		fmt.Println("请输入密码")
		fmt.Scanln(&pwd)
		if name == "张无忌" && pwd == "888" {
			fmt.Println("登录成功")
			break
		}
		fmt.Printf("还有%d次机会", i-1)
	}
}

func test1(a int) {
	fmt.Println("a=", a)
}
