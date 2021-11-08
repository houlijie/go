package main

import "fmt"

/*
	快排 o(nlgn) 不稳定
	// 从序列中选择一个轴点元素， 分割成两个元素
	// 左边< 轴点 右边 > 轴点
	// 对左右节点进行快排

	// 本质 逐渐将每一个元素变成轴点元素
 */

func dqsort(a []int, begin, end int) []int  {
	// 确定轴点
	if begin > end {
		return a
	}

	middle := pivotIdx(0, len(a))
	qsort(a, 0, middle)
	qsort(a, middle, end)
}

func main() {
	arr := []int{3, 100, 10,32, 3, 223, 21, 32}
	// dqsort(arr)

	qs(arr)

	fmt.Println(arr)
}

func qs(a []int) []int  {
	if len(a) < 2 {
		return a
	}

	pivot := a[0]

}
