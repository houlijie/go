package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("conn err ....")
		return
	}

	fmt.Printf("conn=%v\n", conn)
	// 发送一行数据
	reader := bufio.NewReader(os.Stdin) // 标准输入
	// 从终端读入一行
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read err...")
	}
	fmt.Println("line=", line)
	bytes, err := conn.Write([]byte(line)) // string转byte
	fmt.Println("发送字节%d", bytes)

}
