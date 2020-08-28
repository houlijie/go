package main

import "fmt"

func orderSearch(arr [5]int, num int) int {
	index := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			index = 1
		}

		if index != -1 {
			break
		}
	}

	return index
}

func binarySort() {

}

func main() {
	var arr [5]int = [5]int{1, 23, 122, 19}
	var num int
	fmt.Println("请输入要查找的数字")
	fmt.Scanln(&num)
	index := orderSearch(arr, num)

	fmt.Println("index=", index)
}
