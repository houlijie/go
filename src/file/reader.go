package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
	ioReader()
	// buffoReader()
}

// 一次性读取文件到内存
// 适用小文件读取
func ioReader()  {
	str, err:= ioutil.ReadFile("example.txt")

	fmt.Println(string(str), err)
}

// 读取一部分到内存 批量读取
func buffoReader()  {
	f, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("open failed")
	}
	defer f.Close()

	// 有缓冲的
	reader := bufio.NewReader(f)
	for  {
		str, err := reader.ReadString('\n')
		fmt.Println(str, err)
		if err != nil {
			if err == io.EOF {

				fmt.Println("finish")
				break
			}
			fmt.Println("err")
			break
		}

		fmt.Println(str)
	}
}
