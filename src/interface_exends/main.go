package main

import "fmt"

/*
	一、区别
		1、场景不同， 继承只要解决代码的复用性和可维护性，接口是主要是为设计好各种规范方法，让其他自定义类型去实现这些方法
		2、接口比继承更灵活
		3、接口在一定程度上实现代码解耦
	二、联系
		1、实现接口是对继承的一种补充
*/
type fliable interface {
	fly()
}
type monkey struct {
	Name string
}

type littleMonkey struct {
	monkey // 继承 littleMonkey继承了monkey所有的变量和方法 可以直接使用
}

func (m monkey) climb() {
	fmt.Println(m.Name, "can climb")
}

func (m littleMonkey) fly() {
	fmt.Println(m.Name, "can fly")
}

func main() {
	lm := littleMonkey{
		monkey{
			Name: "lm",
		},
	}
	lm.climb()
	lm.fly()

	lm2 := monkey{Name: "lisi"}
	lm2.fly()
}
