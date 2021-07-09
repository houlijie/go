package main

import (
	"fmt"
	_ "sync"
	_ "time"
)

func odd1(x int, A chan bool, B chan bool)  {
	for i := 1; i <= x; i++ {
		ok := <- A
		if ok && i % 2 == 1 {
			fmt.Println("i = ", i)
			B <- true
		}
	}
}

func even1(x int, A chan bool, B chan bool, Exit chan bool)  {
	defer func() {
		close(Exit)
	}()
	for i := 1; i <= x; i++ {
		<- B
		if i % 2 == 0 {
			fmt.Println("i = ", i)
			A <- true
		}
	}
}

func main()  {
	A := make(chan bool, 1)
	B := make(chan bool)
	Exit := make(chan bool)
	x := 10
	A <- true
	go odd1(x, A, B)
	go even1(x,  A, B, Exit)
	<-Exit
}
