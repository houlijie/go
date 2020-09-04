package main

import (
	"fmt"
	"reflect"
)

type stu struct {
	Name string
}

func reflecttest1(i interface{}) {
	// 获取reflect.value
	rv := reflect.ValueOf(i)
	rt := reflect.TypeOf(i)
	fmt.Printf("rv=%v, rt=%v, rt type = %T\n", rv, rt.Kind(), rv) //rv={haha}, rt=main.stu

	// rv转成interface{}
	rvi := rv.Interface()
	fmt.Printf("rvi=%v, rvi type = %T\n", rvi, rvi) //rvi={haha}, rvi type = main.stu

	// 类型断言转成对应的类型
	s1, ok := rvi.(stu)
	if ok {
		fmt.Println(s1.Name)
	}
}
func main() {
	s1 := stu{Name: "haha"}
	reflecttest1(s1)
}
