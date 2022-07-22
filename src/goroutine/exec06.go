package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
var l sync.RWMutex

var count = 20
func main()  {
	wg.Add(1)
	l.Lock()
	go t("窗口1")
	go t("窗口2")
	go t("窗口3")
	go t("窗口4")

	time.Sleep(10 * time.Second)
}

func t(name string)  {
	for {
		if count > 0 {
			fmt.Println(name, ":售出", count)
			count--
		} else {
			fmt.Println(name,"售罄")
			break
		}
	}
}
