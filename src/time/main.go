package main

import (
	"fmt"
)

func a(s []int)  {
	s = []int{6,6,6}
}

func b(s []int)  {
	s[1] = 100
}

func main()  {
	s := make([]int, 1)
	if a := 1; false {
	} else if b := 2; false {
	} else {
		println(a, b)
	}
	var m map[string]int
	m["a"] = 1
	if v, ok := m["b"]; ok {
		fmt.Println(v)
	}

	s := []int{5}
	s = append(s, 7)

	s = append(s, 9)
	fmt.Println(cap(s))
	x := append(s, 11)
	y := append(s, []int{14,15}...)

	fmt.Println(s, cap(s), x, y)
	return
	// s := []int{1,2,3}
	// a(s)
	// fmt.Println(s)
	// b(s)
	//
	// fmt.Println(s)

	return
	// a := []int{1,2,3,4,5,6}
	//
	// i := 2
	//
	// x := a[:0]
	//
	// for index, v := range a {
	// 	if index != i {
	// 		x = append(x, v)
	// 	}
	// }

	// fmt.Println(len(x), x, cap(x), cap(a))
	// t1 := time.Now()
	//
	// fmt.Printf("%T\n", t1)

}
