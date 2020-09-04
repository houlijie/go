package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	map1 = make(map[int]int, 100)
	// lock 全局互斥锁
	lock sync.Mutex
)

func test(n int) {
	x := 1
	for i := 1; i <= n; i++ {
		x = x * i
	}
	lock.Lock()
	map1[n] = x
	lock.Unlock()
}

func main() {
	for i := 0; i < 100; i++ {
		go test(i)
	}

	time.Sleep(10 * time.Second)

	lock.Lock()
	for n, v := range map1 {
		fmt.Printf("n=%d, v=%d\n", n, v)
	}
	lock.Unlock()
}
