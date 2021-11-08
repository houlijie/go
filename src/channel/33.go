package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

func main()  {

	timeStr := "20210607"
	timeDateStr,_ := time.ParseInLocation("20060102", timeStr, time.Local)
	for i:=1;i<=7;i++ {
		fmt.Println(i)
		date := timeDateStr.AddDate(0, 0, i).Format("20060101")
		fmt.Printf("date is %v\n", date)
	}

	return


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