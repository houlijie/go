package main

import (
	"flag"
	"fmt"
	"os"
)

func main()  {
	flagArgs()
-u}

// 实现了命令行参数的解析
// eg mysql -u root -p 127.0.0.1 -port xxx
func flagArgs()  {
	var user, password, host string
	var port int

	flag.StringVar(&user, "u", "默认值", "默认为默认值")
	flag.StringVar(&password, "p", "", "默认主机为默认值")
	flag.StringVar(&host, "h", "", "默认为默认值")
	flag.IntVar(&port, "port", 3306, "port")

	flag.Parse()

	fmt.Println(user, password, host, port)
}

func args()  {
	// 获取命令行参数
	fmt.Println("len of args: ", len(os.Args))
	for _, v := range os.Args {
		fmt.Println("v := ", v)
	}
}