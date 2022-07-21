package main

import (
	"fmt"
	_ "strconv"
	"unsafe"
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

type hchan struct {
	// channel中元素的个数
	qcount uint
	// 缓冲区环形队列大小
	dataqsize uint
	// 缓冲区数组指针
	buf      unsafe.Pointer
	elemsize uint16

	closed uint32
	// 元素类型
	elemtype *_type
	// 发送操作的位置index
	sendx uint
	// 接受操作的index
	recvx uint
	// 待接收列表
	recvq waitq
	// 待发送列表
	sendq waitq
}

//  sudog展示了一个在等待列表中的g， 比如说待发送/待接收
type sudog struct {
	g    *g
	next *sudog
	prev *sudog
	elem unsafe.Pointer
}
type waitq struct {
	first *sudog
	last  *sudog
}

func main() {

	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			fmt.Println(i)
		case 1:
		case 2:
			fmt.Println(2)
		}

	}
	return
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
