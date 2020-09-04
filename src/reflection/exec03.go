package main

import (
	"fmt"
	"reflect"
)

/*
	反射的实践
	1、利用反射来遍历结构体字段， 调用结构体的方法，并获取结构体的标签的值
*/

type mon struct {
	Name string `json:"name"`
}

func (m mon) Print() {
	fmt.Println("call ...")
}

func (m mon) Add(x int, y int) int {
	sum := x + y
	fmt.Println("add ...", sum)
	return sum
}

func test03(i interface{}) {
	iv := reflect.ValueOf(i)
	it := reflect.TypeOf(i)
	num1 := iv.NumField()
	for i := 0; i < num1; i++ {
		// 获取标签
		tagVal := it.Field(i).Tag.Get("json") // name
		if tagVal != "" {
			fmt.Println(tagVal)
		}
	}

	//  获取方法, 方法排序按照asc码大小排序
	numMethod := iv.NumMethod()
	fmt.Println("numMethod=", numMethod)

	var params []reflect.Value
	// 写法1
	// params = append(params, reflect.ValueOf(10))
	// params = append(params, reflect.ValueOf(20))
	// 写法2
	params = make([]reflect.Value, 2)
	params[0] = reflect.ValueOf(10)
	params[1] = reflect.ValueOf(10)

	num44 := iv.Method(0).Call(params) // 有参数
	fmt.Println("num44=", num44[0])

	iv.Method(1).Call(nil) // 无参数
}
func main() {
	m1 := mon{Name: "33"}
	test03(m1)
}
