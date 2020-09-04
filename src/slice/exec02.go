package main

import "fmt"

/*
	//1.不同类型的切片无法复制
	//2.如果s1的长度大于s2的长度，将s2中对应位置上的值替换s1中对应位置的值
	//3.如果s1的长度小于s2的长度，多余的将不做替换
*/
func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s3 := []int{6, 7, 8, 9}
	n1 := copy(s1, s2) // copy返回len(s1), len(s2)较小的那个
	fmt.Println(n1)    // 2
	fmt.Println(s1)    //[4 5 3]
	copy(s2, s3)
	fmt.Println(s2) //[6 7]

	a1 := []int{1, 2}
	var a2 []int = make([]int, 2, 2)
	copy(a2, a1)
	a1 = append(a1, 3)
	fmt.Println(&a1, &a2)
	fmt.Println(&a1, &a2)
}
