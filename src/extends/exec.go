package main

import "fmt"

type mon struct {
	name string
	age  int
}

type E struct {
	mon
	int // 同名结构体不能重复继承，想要进程就要起别名
	n   int
}

func main() {
	a1 := E{mon{name: "sshs", age: 10}, 10, 2}
	fmt.Println(a1)
}
