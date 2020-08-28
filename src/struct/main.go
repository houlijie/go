package main

import "fmt"

type emp struct {
	name string
	age  int
	sex  int
}

type emp1 struct {
	name string
	attr struct {
		name string
	}
}

type person struct {
	name string
	age  int
}

// 指针 slice map的零值都是nil, 需要先make再使用
type emp2 struct {
	name  string
	ptr   *int
	slice []int
	map1  map[string]int
}

func main() {
	a := [...]int{12, 78, 50} // 可变长度的arr 长度是数组的长度 一旦确定不能扩展
	// a[3] = 100                // 这里就会报错 err invalid array index 3 (out of bounds for 3-element array)
	fmt.Println(a)

	// 创建&使用方式
	// 1
	var b1 emp
	fmt.Println("b1 = ", b1) // { 0 0}
	// 2
	var b2 = emp{}
	fmt.Println("b2 = ", b2) // { 0 0}
	// 3
	var b3 *emp = new(emp)
	(*b3).name = "t3"         // 标准写法， 等同于 b3.name = "t3"
	fmt.Println("b3 = ", *b3) // {t3 0 0}
	b3.name = "t31"
	fmt.Println("b3 = ", *b3) // {t31 0 0}
	// 4
	var b4 *emp = &emp{}
	(*b4).name = "b4"
	fmt.Println("b4 = ", *b4) //{b4 0 0}

	// 结构体是值类型
	a1 := emp{
		name: "1111",
		age:  10,
		sex:  1,
	}
	fmt.Println(a1)

	a2 := emp{"jade", 21, 1}
	fmt.Println(a2)

	// 匿名结构体
	a3 := struct {
		name   string
		salary float32
	}{
		name:   "ll",
		salary: 100.33,
	}

	fmt.Println(a3)

	var a4 emp
	a4.name = "hahah"
	fmt.Println(a4) // {hahah 0 0}

	// 结构体是值类型
	a5 := a4
	a5.name = "这是a5"
	fmt.Println(a5) // {这是a5 0 0}
	fmt.Println(a4) // {hahah 0 0}

	var a6 emp1
	a6.name = "emp1"
	a6.attr.name = "hahah"
	fmt.Println(a6)

	var a7 emp2
	a7.name = "test"
	fmt.Println(a7) // {test <nil> [] map[]}
	if a7.slice == nil {
		fmt.Println("slice is nil")
	}
	// a7.slice[0] = 1 //  error: index out of range
	a7.slice = make([]int, 2) // {test <nil> [1 0] map[]}
	a7.slice[0] = 1

	a7.map1 = make(map[string]int)
	a7.map1["tets"] = 1
	fmt.Println(a7) // {test <nil> [1 0] map[tets:1]}

	a8 := a7
	fmt.Println("at=", a7)
	fmt.Println("a8=", a8)
	a7.map1["tets"] = 10
	fmt.Println("a7=", a7)
	fmt.Println("a8=", a8)

	a9 := person{
		name: "zhangsan",
		age:  12,
	}

	var a10 *person = &a9
	a9.name = "lisi"
	fmt.Println(*a10)      // {lisi 12}
	(*a10).name = "wangwu" // 这里注意  .的优先级高于* ， 写法 *a10.name是粗我无的写法
	fmt.Println(a9)        // {lisi 12}
	fmt.Println(*a10)      // {lisi 12}
}
