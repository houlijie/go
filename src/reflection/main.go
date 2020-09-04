package main

import (
	"fmt"
	"reflect"
)

/**
	注意事项
	1、reflect.ValueOf.Kind 等同于reflect.IndexOf.Kind是一个常量
	2、
**/
func reflecttest(i interface{}) {
	rv := reflect.ValueOf(i)
	rt := reflect.TypeOf(i)
	fmt.Printf("rv=%v, rt=%v\n", rv, rt.Kind()) //rv={haha}, rt=main.stu
	var n1 int64 = 10
	// n2 := rv + n1 ////invalid operation: rv + n1 (mismatched types reflect.Value and int)
	n2 := rv.Int() + n1
	fmt.Println("n2=", n2)
}
func main() {
	var s2 int64 = 10
	fmt.Println("s2=", s2)
	reflecttest(s2)
}
