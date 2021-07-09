package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	t1 := time.Now().Nanosecond()
	fileName := "writer.txt"
	f, err := os.Create(fileName)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	newWriter := bufio.NewWriter(f)

	maxSize := int64(4 * 1024 * 1024)
	str := "hello"
	l := int64(len(str))
	count := int64(0)
	for {
		if (count + l + 1) <= maxSize {
			_, err :=  newWriter.WriteString(","+str)
			if err != nil {
			panic(err.Error())
			}
			count = count + l+1
		} else {
			min := maxSize - count
			for i := int64(0); i < min; i++ {
				_, err := newWriter.WriteString("a")
				if err != nil {
					panic(err.Error())
				}
			}
			break
		}
	}
	info, err := f.Stat()
	fmt.Println(info.Size(), time.Now().Nanosecond() - t1)
}