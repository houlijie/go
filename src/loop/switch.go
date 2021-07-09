package main

import (
	"fmt"
)

func main() {
	switch finger := 8; finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default: // 默认情况
		fmt.Println("incorrect finger number")
	}

	var str = "hello"
	switch str {
	case "hellp", "wof": // 一分支多值
		fmt.Println("hello")
	case "hello":
		fmt.Println("this is the crorrect word")
	}

	// fallthrough 穿透，执行结束后可以继续执行
	var r int = 11
	switch {
	case r > 10 && r < 20: // case 可以支持表达式
		fmt.Println(r)
		fallthrough
	case r == 11:
		fmt.Println("this is 11")
	}
}
