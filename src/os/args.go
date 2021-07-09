package main

import (
	"fmt"
	"os"
)

func main()  {
	// 获取命令行参数
	fmt.Println("len of args: ", len(os.Args))
	for _, v := range os.Args {
		fmt.Println("v := ", v)
	}
}