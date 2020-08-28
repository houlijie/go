package main

import "fmt"

func bubble(arr *[5]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}
}
func main() {
	var arr [5]int = [...]int{10, 3, 223, 21, 32}
	bubble(&arr)

	fmt.Println(arr)
}
