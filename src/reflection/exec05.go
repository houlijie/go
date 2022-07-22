package main

import (
	"fmt"
	"sort"
)

func main()  {
	a := []int{7, 3, 2, 6, 8,1,9,5,4, 6, 10, 6}
	// b := sort(a)
	sort.Ints(a)

	fmt.Println(b)

}

// func sort(a []int) []int  {
// 	if len(a) < 2 {
// 		return a
// 	}
//
// 	len := len(a)
// 	left, right := 0, len -1
// 	for left < right {
// 		sort.Ints(a)
// 		if a[left] > a[right] {
// 			tmp := a[left]
// 			a[left] = a[right]
// 			a[right] = tmp
// 			left++
// 		} else {
//
// 		}
// 		fmt.Println(left, right, a[left], a[right],  a[left] <= a[right])
// 		if a[left] <= a[right] {
// 			left++
// 		} else {
// 			fmt.Println("sssss")
// 			right--
// 			fmt.Println(right, a[left], a[right])
// 		}
// 		// fmt.Println(left, right, a[left], a[right])
// 	}
//
// 	return a
// }