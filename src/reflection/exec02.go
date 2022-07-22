package main

import (
	"fmt"
	"reflect"
)

type cat struct {
	Name string
}

// 通过反射修改Name
func test(i interface{}) {
	singlefl
	v := reflect.ValueOf(i) // v kind is ptr
	vt := reflect.TypeOf(i)

	// aa : = v.Elem().FieldByName("Name").setString("hah")
	// v.SetString("katty") //panic: reflect: reflect.Value.SetString using unaddressable value

	fmt.Println(v.Elem().FieldByName("Name"), vt.Elem())

}
func main() {
	c1 := cat{Name: "sha"}
	test(&c1)
}
