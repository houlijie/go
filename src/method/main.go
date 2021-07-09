package main

import (
	"fmt"
)

// 方法和函数的区别
/*
	1、方法作用在指定的的数据类型上 (即： 和指定的数据类型绑定)， 因此自定义类型都可以有方法，而不仅仅是struct
	2、在通过一个变量去调用方法时，其调用机制和函数一致,但是调用方法时， 该变量本身也会作为一个参数传递到方法，类型和原变量保持一致
	3、调用方式不同 函数： 函数名（参数列表） 方法：变量.方法名(参数)
	4、对于普通函数，接受者为值类型时，不能将指针类型的数据直接传递，反之亦然， 即：入参和定义格式必须一致
	5、对于方法，接收者为值类型时，可以直接用指针类型的变量调用方法，反之亦然
*/

// 总结
/*
	1、不管调用形式如何，真正决定是值copy 还是地址copy 是看这个方法是和哪个类型绑定
	2、()

*/
// 声明和调用
type Person struct {
	Name string
	Age  int
}

// 值copy
func (p Person) jisuan(num int) int {
	var sum int
	for i := 0; i <= num; i++ {
		sum += i
	}
	return sum
}

// 地址copy
func (p *Person) test1() {
	p.Name = "haha"
	fmt.Println(p.Name)
}

func (p Person) test() { // 是一个副本，不影响原值
	p.Name = "jade"
	fmt.Println(p.Name)
}

func main() {
	var a1 Person
	a1.Name = "tom"
	a1.test()            // 2： jade
	fmt.Println(a1.Name) //tom
	// 接收者为值类型，可以直接用指针类型的变量调用方法，反过来通用也可以
	(&a1).test() // jade 这个和2 还是值copy, test name 变更不影响a1

	(&a1).test1()        // haha 地址copy
	fmt.Println(a1.Name) // haha

	total := a1.jisuan(2) // 3
	fmt.Println(total)

	// 方式一， 缺点是结构体顺序调整 可能会报错
	b1 := Person{"zhangsan", 10}
	fmt.Println(b1)

	// 方式二
	b2 := Person{
		Age:  20,
		Name: "lisi",
	}
	fmt.Println(b2)

	// 返回结构体的指针类型！！！
	b3 := &Person{"wangwu", 2} // 原理： b3-->地址-->结构体数据{wangwu 2}
	fmt.Println(*b3)

	b4 := b3
	b4.Name = "sha"
	fmt.Println(*b3) // {sha 2}
}
