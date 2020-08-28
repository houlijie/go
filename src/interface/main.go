package main

import "fmt"

// 接口 interface 引用类型 未初始化是nil
/*
	1、方法实现了（*所有）interface接口，就认为实现了interface
	2、只要是自定义类型 都可以实现接口

*/

type usb interface {
	// name string // error 不能包含任何变量
	start() // 方法不需要实现
	stop()
}

type usb1 interface {
	end()
	// start() // error: duplicate method start
}

// 接口继承多个, 如果要实现usb2 必须实现所有方法(start\stop\end), 缺一不可
// 多重继承时 方法名不能重复
type usb2 interface {
	usb
	usb1
}

type phone struct {
}

func (p *phone) start() {
	fmt.Println("phone start....")
}

func (p *phone) stop() {
	fmt.Println("phone stop....")
}

type camera struct {
}

func (p *camera) start() {
	fmt.Println("camera start....")
}

func (p *camera) stop() {
	fmt.Println("camera stop....")
}

func (p *camera) end() {
	fmt.Println("camera end....")
}

type computer struct {
}

func (c computer) exec(u usb) {
	u.start()
	u.stop()
}

func main() {
	com := computer{}
	ph := phone{}
	cam := camera{}

	com.exec(&ph)
	com.exec(&cam)

	// 多重继承
	var b2 usb2 = &cam
	b2.end()
}
