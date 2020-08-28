package main

import "fmt"

type Circle struct {
	radius float64
}

func (c *Circle) String() string {
	return "sssss"
}

func (c Circle) area() float64 {
	return 3.14 * (c.radius) * c.radius
}

func (c *Circle) test() {
	(*c).radius = 2
	fmt.Println(*c)
}

func main() {
	c := Circle{3}
	fmt.Println(&c) //输出sssss 默认会调用这个变量的String() string{return xx}进行输出
	num := c.area()
	fmt.Println(num)

	(&c).test()

}
