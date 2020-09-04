package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	chi := make(chan interface{}, 4)

	a1 := 10
	a2 := Cat{Name: "zhansan", Age: 10}
	a3 := 12.34
	a4 := "jack"

	chi <- a1
	chi <- a2
	chi <- a3
	chi <- a4

	r1 := <-chi
	r2 := <-chi
	// fmt.Println(r2.Name) // errorr2.Name undefined (type interface {} is interface with no methods) 接口是一般类型，不知道具体类型， 需要使用类型断言
	fmt.Println(r1, r2.(Cat).Name)
}
