package main

import (
	"fmt"
)

func main() {
	// 使用方式一：直接引用一个数组
	// 切片是数组的一个引用， 遵循引用传递的机制
	// 切片的长度是可变的， 是一个可以动态变化数组
	var arr [3]int = [...]int{11, 222, 3222} // 数组
	// slice1 引用arr这个数组  包含1 不包含3
	slice1 := arr[1:2]
	fmt.Println("slice1 长度是", len(slice1)) // 1
	fmt.Println("slice1 元素是", slice1)      //[222]

	// fmt.Println("slice1[2]是", slice1[2]) //runtime error: index out of range 不能越界
	// slice1[2] = 1000//runtime error: index out of range 不能越界
	slice1 = append(slice1, 22, 333)  // 追加元素
	fmt.Println("slice1追加后是", slice1) // [222 22 333]
	// 切片的容量, 是动态的
	fmt.Println("slice1 容量是", cap(slice1)) // 2

	arr[1] = 1999
	fmt.Println("slice1 元素是", slice1) //[1999] 因为是引用

	// 使用方式二
	// 使用内置函数make， make底层创建一个数组，由切片在底层维护
	var slice2 []int = make([]int, 4, 10) // make(类型, len, cap)
	fmt.Println(slice2)

	sliceAppend := append(slice1, slice1)
	fmt.Println("sliceAppend=", sliceAppend)

	var slice3 []int
	// slice3[0] = 11 // 这里就是错误的， runtime error: index out of range
	fmt.Println(slice3) // []

	// 使用方式三
	var slice4 []string = []string{"jack", "lj"}
	fmt.Println(slice4)      //[jack lj]
	fmt.Println(len(slice4)) // 2
	fmt.Println(cap(slice4)) // 2

	// 切片可以继续切片, 指向的数据空间是同一个
	slice5 := slice4[1:2]
	fmt.Println(slice5) // [lj]

	// 遍历
	fmt.Println(slice1)
	for i := 0; i < len(slice1); i++ {
		fmt.Println(i)
		fmt.Println(slice1[i])
	}
}
