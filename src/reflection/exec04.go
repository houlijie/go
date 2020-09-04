package main

import (
	"fmt"
	"reflect"
	"time"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (c Cal) GetSub(name string) {
	res := c.Num1 - c.Num2
	fmt.Printf("%v 完成了减法：%d - %d = %d", name, c.Num1, c.Num2, res)
}

func test04(i interface{}) {
	val := reflect.ValueOf(i)
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf("tom")
	val.MethodByName("GetSub").Call(params)
}

func main() {
	n1 := Cal{10, 1}
	test04(&n1)

	time.Sleep(time.Second * 10)
}
