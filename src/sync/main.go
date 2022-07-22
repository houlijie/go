package main

import (
	"fmt"
	"sync"
	"time"
)

var rwl *sync.RWMutex
var o sync.Once
func main()  {
	rwl = new(sync.RWMutex)
	go t(1)
	go t(2)
	o.Do()

	time.Sleep(2 * time.Second)
}

func t(i int) {
	fmt.Println(i, " start...")
	rwl.RLock()
	fmt.Println(i, " reading...")
	time.Sleep(time.Microsecond * 10)
	rwl.RUnlock()
	fmt.Println(i, " read unlock...")
}
