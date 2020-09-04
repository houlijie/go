package main

import "fmt"

func main() {
	n1 := 1
	select {
	case n1 != 2:
		fmt.Println("n1 != 2")
	}
}
