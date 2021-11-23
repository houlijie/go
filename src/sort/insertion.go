package main

import "fmt"

/**
	插入排序 o(n^2) 稳定
**/

func insertionS(a []int) []int {
	len := len(a)
	if len <= 1 {
		return a
	}

	for i := 1; i < len; i++ {
		tmp := a[i]
		j := i -1
		for ; j >= 0 && a[j] > tmp; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = tmp
	}

	return a
}

func main()  {
	arr := []int{3, 100, 10, 3, 223, 21, 32}
	fmt.Println(insertionS(arr))
}
