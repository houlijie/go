package main

import (
	"factorymodel/model"
	"fmt"
)

// 工厂模式
/*
	1、场景：解决跨包创建结构体实例
		如果model包的结构体变量首字母是大写， 引入后可以直接使用， 如果小写 则就用到工厂模式
*/

func main() {
	a1 := model.NewStudent("tom", 90) // 指针
	fmt.Println(*a1)
}
