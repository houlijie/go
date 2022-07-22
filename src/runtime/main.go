package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
	}()

	for i := 5; i < 10; i++ {
		runtime.Gosched()
		fmt.Println(i)

	}

	// go func() {
	// 	fmt.Println("start t ....")
	// 	t()
	// 	fmt.Println("end t ....")
	//
	// }()
	//
	// time.Sleep(time.Second * 3)
}

func t()  {
	defer fmt.Println("defer...")
	// runtime.Goexit() // 不会打印processing 和 end， 会执行satrt-->defer
	// return // 不会打印processing， 会执行start-->defer-->end

	fmt.Println("processing t....")
}


