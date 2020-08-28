package main

import (
	"encoding/json"
	"fmt"
)

type stu1 struct {
	name string
	age  int
}

type stu3 struct {
	Name string
	age  int
}

type stu4 struct {
	Name string `json:"name"` // 这个就是结构体标签tag，tag可以通过反射机制获取， 常见的场景是序列化和反序列化
	age  int
}

type stu2 stu1

type integer int

func main() {
	var a1 = stu1{}
	var a2 = stu2{}

	// a2 = a1 // error cannot use a1 (type stu1) as type stu2 in assignment
	a2 = stu2(a1)
	fmt.Println(a1, a2)

	var num1 integer = 10
	var num2 int = 10

	// 注意** 这里报错是: 结构体新定义int alias interger， golang认为是新的数据类型 但是可以强制互相转换
	// num2 = num1 // cannot use num1 (type integer) as type int in assignment
	// num2 = int(num1) // 强制转换成int
	// num1 = integer(num2) // 强制转换成integer
	fmt.Println(num1, num2)

	a4 := stu1{"zhangsan", 10}
	fmt.Println(a4)

	jsona4, _ := json.Marshal(a4) // {}  这里是因为Marshal相当于在其他包内访问stu1, 首字母小写不能在其他包使用
	fmt.Println(string(jsona4))

	a5 := stu3{"lisi", 20}
	jsona5, _ := json.Marshal(a5)
	fmt.Println(string(jsona5)) // { "Name":"lisi"} 注意这里的大写

	a6 := stu4{"wangwu", 20}
	jsona6, _ := json.Marshal(a6)
	fmt.Println(string(jsona6)) // {"name":"wangwu"}  name 是小写了，json tag起作用
}
