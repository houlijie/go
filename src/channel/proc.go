package main

import (
	"fmt"
	"time"
)

// 要求每秒钟调用一次proc并保证程序不退出?
func main() {
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出
		t := time.NewTicker(time.Second)
		defer t.Stop()
		for  {
			select {
			case <-t.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					proc()
				}()
			}
		}
	}()

	select {
	}
}

func proc() {
	panic("ok")
}
