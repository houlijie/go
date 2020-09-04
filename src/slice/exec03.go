package main

import "fmt"

func main() {
	skus := []int{1, 2, 3, 4, 5}
	ret := make(map[uint64][]int, 1, 1)
	ret[1] = 100
	for _, item := range skus {
		list, ok := ret[1]
		if !ok {
			list = make([]int, 0)
			ret[1] = list
		}
		list = append(list, item)
		// ret[1] = list
	}
	fmt.Println(ret)
}
