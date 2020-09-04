package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		// 创建新的切片
		buf := make([]byte, 1024)
		// 1、等待客户端通过conn发送信息
		// 2、如果客户端没有write 协程就一直阻塞
		bytes, err := conn.Read(buf)
		if err != nil {
			fmt.Println("server read err: ", err) // EOF 客户端已退出
			return
		}
		// 3. 显示客户端发送的内容到服务器的终端
		fmt.Printf(string(buf[:bytes])) // 显示读到的字节
	}
}

func main() {
	fmt.Println("start...")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err...")
		return
	}
	defer listen.Close()

	for {
		// 等待客户端连接
		fmt.Println("conn waiting.....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("conn err.....")
		}
		fmt.Printf("conn succ\n conn=%v\nclient ip = %v\n", conn, conn.RemoteAddr().String())

		// 启动协程， 为客户端服务
		go process(conn)
	}

	fmt.Printf("listen suc= %v", listen)
}
