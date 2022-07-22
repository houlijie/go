package main

import "fmt"

/*
	二分查找

*/
func main()  {
	a := []int{1, 2, 4, 6, 18, 29, 30}
	fmt.Println(search(a, 1))
}

func search(arr []int, n int) int  {
	len := len(arr)
	l, r := 0, len - 1
	for l <= r {
		mid := (l + r) / 2 // 这里l+r可能溢出， 可以优化为
		fmt.Println(l, r, mid, arr[mid])
		if arr[mid] == n {
			return mid
		}

		if arr[mid] < n {
			l = mid + 1 // 这里必须要+-1， 不然可能会存在死循环的场景
		} else {
			r = mid - 1
		}
	}

	return -1
}
