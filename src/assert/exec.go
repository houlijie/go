package main

import "fmt"

// 断言实例1
type sub interface {
	open()
	stop()
}

type Camera struct {
	name string
}

type Computer struct {
	name string
}

func (c Camera) open() {
	fmt.Println(c.name, "is open...")
}

func (c Camera) stop() {
	fmt.Println("camera is stop...")
}

func (c Camera) working() {
	fmt.Println("computer is working...")
}

func (c Computer) open() {
	fmt.Println(c.name, "is open.....")
}

func (c Computer) stop() {
	fmt.Println("computer is stop...")
}

func (c Computer) exec(s sub) {
	s.open()
	camera, ok := s.(Camera)
	if ok {
		camera.working()
	}
	s.stop()
}

func main() {
	var compouter Computer
	sub1 := [2]sub{Camera{name: "cam"}, Computer{name: "comp"}}
	for _, v := range sub1 {
		compouter.exec(v)
	}
}
