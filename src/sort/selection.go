package main

import (
	"fmt"
)

/*
	选择排序 (不稳定) o(n^2), o(n^2)
	//  从数列中找到最大， 跟最末尾交换位置
 */
func sort(a []int) []int  {
	if len(a) == 0 {
		return a
	}

	for end := len(a) - 1; end > 0 ; end-- {
		maxIndex := 0
		for i := 0; i <= end; i++ {
			if a[maxIndex] < a[i] {
				maxIndex = i
			}
		}
		a[maxIndex], a[end] = a[end], a[maxIndex]
	}

	return a
}

func s(a []int) []int   {
	for end := len(a) - 1 ; end > 0 ; end -- {
		maxIdx := 0
		for i := 1; i <= end; i++ {
			if a[maxIdx] <= a[i] {
				maxIdx = i
			}
		}

		a[maxIdx], a[end] = a[end], a[maxIdx]
		fmt.Println(a)
	}

	return a
}

func main() {
	arr := []int{3, 100, 10,32, 3, 223, 32, 21, 32}
	sort(arr)

	// s(arr)

	fmt.Println(arr)
}

