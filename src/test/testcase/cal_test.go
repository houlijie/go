package testcase

import (
	"fmt"
	"testing"
)

/*
 1、全部
	go test
 2、 输出日志
	gp test -v
 2、测试单个文件 一定要带上被测试的源文件
 	go test -v cal_test.go cal.go
 3、测试指定方法
	go test -v -test.run TestAdd
*/
func TestAdd(t *testing.T) {
	res := add(10, 11)
	fmt.Println(res)
}

func TestMin(t *testing.T) {
	res := minus(10, 11)
	fmt.Println(res)
}
