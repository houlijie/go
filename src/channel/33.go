package main

import (
	"bytes"
	"os"
)

func main()  {
	str := "hello"

	f, err := os.OpenFile("hlj2.log", os.O_CREATE, 0775)
	if err != nil {
		panic("打开文件失败")
	}

	maxSize := 1024 * 1024 * 4

	strLen := len(str)

	bytes.Buffer


	defer f.Close()

}