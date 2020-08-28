package main

import (
	"fmt"
	"strconv"
)

// 多重循环
func main() {
	var str1 = "hello"
	var str2 bool = true

	str2, _ = strconv.ParseBool(str1) // 返回false 如果转不成功， 会返回默认值. 而不是初始值
	fmt.Println("str=", str2)

	var str string = "哈哈哈"
	fmt.Println("str=", str)

	// var int1 int32 = 23444
	// var int2 int8
	// var int3 int64

	// int2 = int1 + 20
	// int3 = int1 - 100
	// fmt.Printf("int2 = %v, int3 = %c", int2, int3)

	// 统计3个班级的几个人数
	var passNum int
	var classNum int = 1
	var studentNum int = 1
	for i := 1; i <= classNum; i++ {
		for j := 1; j <= studentNum; j++ {
			var score float64
			fmt.Printf("请输入第%d个学生的成绩 \n", i)
			fmt.Scanln(&score)
			if score >= 60 {
				passNum += 1
			}
		}
	}

	// 统计3个班的平均分， 每个班有5个同学， 成绩从键盘输入
	total := 0.0
	for j := 1; j <= classNum; j++ {
		sum := 0.0
		for i := 1; i <= studentNum; i++ {
			var score float64
			fmt.Printf("请输入第%d个学生的成绩 \n", i)
			fmt.Scanln(&score)
			sum += score
		}
		total += sum

		fmt.Printf("第%d班级的平均分是%v \n", j, sum/float64(studentNum))
	}

	fmt.Printf("班级的平均分是%v", total/float64(studentNum*classNum))
}
