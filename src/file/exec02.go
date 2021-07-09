package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 统计文件内英文、数字、空格、其他字符的个数
func main()  {
	f, err := os.Open("copy.txt")
	if err != nil {
		fmt.Println("open failed", err)
	}
	defer f.Close()

	var numCount, charCount, spaceCount uint64

	reader := bufio.NewReader(f)

	for  {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, s := range str {
			fmt.Printf("zzz:%c", s)
			if s >= 'a' && s <= 'z' {
				charCount++
			}

			if s >= 'A' && s <= 'Z' {
				charCount++
			}

			if s >= '0' && s <= '9' {
				numCount++
			}

			if s == ' ' || s == '\t' {
				spaceCount++
			}
		}
	}

	fmt.Println(charCount, numCount, spaceCount)
}
