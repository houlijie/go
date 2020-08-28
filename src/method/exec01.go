package main

import "fmt"

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (s Student) say() string {
	return fmt.Sprintf("name=%v, gender=%v, age=%d, id =%d, score=%d", s.name, s.gender, s.age, s.id, s.score)
}

type Visitor struct {
	name string
	age  int8
}

func (v *Visitor) check() {
	if v.age > 18 {
		fmt.Println("免票")
	} else {
		fmt.Println("也免票")
	}
}

func main() {
	a1 := Student{"zhangsan", "nv", 10, 111, 90.8}
	fmt.Println(a1.say())

	a2 := Visitor{"lis", 10}
	(&a2).check() // 等同于a2.check() //底层兼容
}
