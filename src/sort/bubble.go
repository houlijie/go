package main

import "fmt"

/*
	冒泡排序(稳定排序)
 */
func bubble(arr []int) {
	for i := 0; i < len(arr); i++ {
		sorted := true
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
				sorted = false
			}
		}

		if sorted {
			break
		}
	}
}

func bubble2(arr []int)  {
	for i := len(arr) - 1; i > 0; i-- {
		sortedIdx := 1
		for j := 1; j <= i; j++ {
			if arr[j] < arr[j - 1] {
				arr[j], arr[j - 1] = arr[j - 1], arr[j]
				sortedIdx = j
			}

		}
		i = sortedIdx
	}
}

func bb(a []int) []int  {
	for i := len(a) - 1; i > 0; i-- {
		sortedIndex := 1
		for j := 1; j < i; j++ {
			if a[j] < a[j - 1] {
				a[j], a[j - 1] = a[j - 1], a[j]
				sortedIndex = j
			}
		}
		i = sortedIndex
	}

	return a
}

func main() {
	arr := []int{3, 100, 10, 3, 223, 21, 32}
	// bubble2(arr)
	bb(arr)

	fmt.Println(arr)
}
