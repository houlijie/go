package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var flag int32

func main() {
	atomic.StoreInt32(&flag, 1)
	go even()
	go odd()
	time.Sleep(time.Second * 21)
}

func even() {
	i := 0
	for i <= 20 {
		v := atomic.LoadInt32(&flag)
		if v == 1 {
			fmt.Println("even =", i)
			atomic.StoreInt32(&flag, 2)
			i = i + 2
			time.Sleep(time.Second * 1)
		}
	}
}

func odd() {
	i := 1
	for i <= 20 {
		v := atomic.LoadInt32(&flag)
		if v == 2 {
			fmt.Println("odd  =", i)
			atomic.StoreInt32(&flag, 1)
			i = i + 2
			time.Sleep(time.Second * 1)
		}
	}
}