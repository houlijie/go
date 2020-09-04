package main

import "fmt"

/*
	指针
	1、基本数据类型， 变量存的是值，
*/
func main() {
	// 方式1
	a1 := 1
	var ptr *int = &a1           // ptr存的是a1的地址， ptr也有自己的地址， prt的值指向的是a1的值
	fmt.Println(&a1)             // 0xc00001a0b0
	fmt.Println(ptr, *ptr, &ptr) // 0xc00001a0b0 1 0xc00000e028

	*ptr = 1000
	fmt.Println(a1, *ptr) // 1000 1000

	a2 := 200
	ptr = &a2
	fmt.Println(a1, a2, *ptr) // 1000 200 200

	// 方式2

	var b1 *int
	fmt.Println(b1, &b1) // <nil> 0xc00000e038
	// *b1 = 100 //不能对nil执行赋值操作
	// fmt.Println(b1) //panic: runtime error: invalid memory address or nil pointer dereference
	/*
		new
		内建函数分配内存。
		参数：基本数据类型，而非值，
		返回值：指向该类型的新分配的零值的指针。

		make
		内建函数分配并初始化一个类型为切片、映射、或（仅仅为）信道的对象。
		与 new 相同的是，参数为类型，而非值。
		不同的是，make 的返回类型 与其参数相同，而非指向它的指针。其具体结果取决于具体的类型：
			切片：size 指定了其长度。该切片的容量等于其长度。第二个整数实参可用来指定
				不同的容量；它必须不小于其长度，因此 make([]int, 0, 10) 会分配一个长度为0，
				容量为10的切片。
			映射：初始分配的创建取决于 size，但产生的映射长度为0。size 可以省略，这种情况下
				就会分配一个小的起始大小。
			信道：信道的缓存根据指定的缓存容量初始化。若 size 为零或被省略，该信道即为无缓存的。
	*/
	var b2 = new(int)
	*b2 = 100
	fmt.Println(b2, *b2) //0xc0000b4040 100
}
