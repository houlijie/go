package main

import (
	"container/list"
	"container/ring"
	"fmt"
)

func main()  {
	a := list.New()
	a.PushFront(10)
	a.PushBack("333")
	a.PushBack("444")

	for e := a.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value, e.Prev())
	}

	fmt.Println(a.Front())

	r := ring.New(1)
	r.Len()

}
