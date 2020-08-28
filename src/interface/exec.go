package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []int{1, 2, 3, 10, 8}
	fmt.Println(slice1)
	(sort.IntSlice(slice1)).Sort()                 // [1 2 3 8 10] 升序排列。
	sort.Sort(sort.Reverse(sort.IntSlice(slice1))) // [10 8 3 2 1] 降序
	fmt.Println(slice1)
}
