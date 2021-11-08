package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main()  {
	a := []int{1, 2, 5}

	t := 3

	m := make(map[int]int, 3)
	for i, v := range a {
		m[v] = i
	}

	for _, v := range a {
		b := t - v
		if i, ok := m[b]; ok {
			fmt.Println(i)
		}
	}

	return
	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan int)

	// 读
	go func() {
		for x := range ch {
			fmt.Println("x=", x)
		}
		defer wg.Done()
	}()

	// 写
	go func() {
		rand.Seed(time.Now().UnixNano())

		for i := 0; i < 5; i++ {
			x := rand.Intn(5)+1
			ch<-x
		}
		defer wg.Done()
		close(ch)
	}()

	wg.Wait()
}
