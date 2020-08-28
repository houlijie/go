package main

import "fmt"

type pupil struct {
	student
	name string
}

type graduate struct {
	student // 匿名函数体, 匿名函数的字段和方法 graduate 都可以使用
}

type student struct {
	name string
	age  int
}

type worker struct {
	name string
	age  int
}

type teacher struct {
	name   string
	age    int
	gender string
}

// 多重继承，不建议使用多重继承
type person struct {
	student // 匿名结构体 嵌入多个拥有相同字段和方法的结构体时，在访问时必须明确指定匿名结构体名字
	worker
	tc teacher // 组合，有名结构体，调用时必须指定结构体名称
}

type personx struct {
	*student
	*worker
	tc *teacher
}

func (s *student) getName() string {
	return s.name
}

func (s *student) getAge() int {
	return s.age
}

func (p *pupil) getName() string {
	return p.name
}

// 继承
func main() {
	// 结构体嵌入匿名函数调用方法
	p1 := &pupil{}
	p1.name = "wangwu"
	p1.student.name = "zhangsan" // 等同 p1.name = "lishi"
	p1.student.age = 10
	fmt.Println(p1.getName())         // wangwu 当前结构体方法优先级高于匿名函数 先找p1 再找student
	fmt.Println(p1.student.getName()) // zhangsan
	fmt.Println(p1.getAge())

	// 多个匿名函数体
	person1 := &person{}
	// person1.name = "zhansa" // error ambiguous selector person1.name
	person1.worker.name = "zhansa" // 必须指定匿名结构体名称

	// person1.gender = "man" // error person1.gender undefined (type *person has no field or method gender)
	person1.tc.gender = "man" // 组合， 必须指定结构体名称
	fmt.Println(*person1)

	// 写法三
	person2 := person{student{"ss", 10}, worker{"www", 10}, teacher{"lis", 0, "man"}}
	fmt.Println(person2)

	person3 := personx{&student{"ss", 10}, &worker{"www", 10}, &teacher{"lis", 0, "man"}}
	fmt.Println(person3)          // {0xc000096000 0xc000096020 0xc000098000}
	fmt.Println(*person3.student) // {ss 10}
}
