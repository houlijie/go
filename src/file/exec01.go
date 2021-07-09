package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main()  {
	copy()
	// rf, err := ioutil.ReadFile("t.txt")
	// if err != nil {
	// 	fmt.Println("read fail", err)
	// 	return
	// }
	//
	// fmt.Println(string(rf))
	// err = ioutil.WriteFile("example.txt", rf, os.ModeAppend)
	// if err != nil {
	// 	fmt.Println("write fail")
	// }
}

func copy()  {
	srcf, err := os.Open("t.txt")
	if err != nil {
		fmt.Println("open src failed", err)
		return
	}
	defer srcf.Close()

	reader := bufio.NewReader(srcf)

	dstf, err := os.OpenFile("copy.txt", os.O_WRONLY|os.O_CREATE, 0755) // 权限很重要
	if err != nil {
		fmt.Println("open dst failed", err)
	}
	defer dstf.Close()
	writer := bufio.NewWriter(dstf)

	x, err := io.Copy(writer, reader)
	fmt.Println(x, err)
	if err != nil {
		fmt.Println("copy filed", err)
	}
}
