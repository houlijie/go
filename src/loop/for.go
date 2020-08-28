package main

import (
	"fmt"
)

func main2() {
	//for循环4要素
	// 1 循环变量初始化
	// 2 循环条件， 一定是返回bool类型的表达式
	// 3 循环语句
	// 4 变量迭代

	// 标准写法
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 循环变量初始化和变量迭代可以放在别的位置
	j := 1
	for j < 10 {
		fmt.Println("j:", j)
		j++
	}

	// 第三种写法, 等同于for ;; {} 无限循环， 这种写法通常需要和break一起使用
	k := 1
	for { // 等价 for ;; {}
		if k < 10 {
			fmt.Println("ok")
			k++
		} else {
			break //break 退出循环
		}
	}

	// 循环判断条件必须是返回一个bool的表达式
	// for i := 0; 1999; i++ { // 会报错：  non-bool 1999 (type int) used as for condition
	// 	fmt.Println(333)
	// }

	// --------------------字符串遍历------------------------
	// 方法1
	var str string = "hello world， 哈哈哈"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
		// fmt.Println(str[i])
	}

	// 注：如果字符串包含中文
	// 字符串遍历方法1会出错， 会出现乱码， 原因是传统的遍历是按照字节来遍历的， 而汉字在utf8编码下对应3个字节
	// 解决方案：需要将str转成[]rune切片
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c \n", str2[i])
		// fmt.Println(str[i])

	}

	// for-range 遍历字符串, 推荐使用, 默认是按照字符的方式遍历，所以中文也不会产生乱码
	for index, val := range str {
		fmt.Printf("index=%d, val=%c \n", index, val)
	}

}
