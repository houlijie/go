package main

// 顺序栈
type ArrayStack struct {
	items []int
	cap uint64
	len uint64
}

func pop() int  {

}

func push(items []int, n int) []int {
	return append(items, n)
}
