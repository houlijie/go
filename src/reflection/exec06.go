package main

import (
	"fmt"
	"reflect"
)

type Stu struct {
	Name string
	Age int
	School string
}

func main() {
	var num float64 = 1.23

	p := reflect.ValueOf(&num)
	newValue := p.Elem()// 如果p不是指针 会panic

	fmt.Println(newValue.Type())
	fmt.Println("是否可修改数据", newValue.CanSet())

	// newValue.SetInt(2) // panic:  reflect: call of reflect.Value.SetInt on float64 Value
	newValue.SetFloat(3.4)
	fmt.Println(num)

	// 结构体改变数值
	s1 := Stu{
		Name: "zhangsan",
		Age: 10,
		School: "hahah",
	}

	fmt.Println(s1)

	// 改变树值
	value := reflect.ValueOf(&s1)
	newV := value.Elem()
	// newV.FieldByName("name").SetString("aaa") // panic: reflect: call of reflect.Value.SetString on zero Value
	newV.FieldByName("Name").SetString("aaa")
	fmt.Println(s1)


}
