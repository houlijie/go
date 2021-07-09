package main

import (
	"bufio"
	"os"
)

func main() {
	filePath := "t.txt"
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}

	defer f.Close()

	str := "hello world\n"
	// 1、普通写入
	// f.WriteString(str)

	// 2、缓存写入
	writer := bufio.NewWriter(f)

	writer.WriteString(str)

	writer.Flush()

}

