package main

import (
	"errors"
	"fmt"
)

type queue struct {
	maxSize int
	array   [4]int
	front   int // 指向队首，不包含第一个元素
	rear    int
}

func (q *queue) AddQueue(val int) (err error) {
	if q.rear == q.maxSize-1 {
		return errors.New("queue full")
	}

	q.rear++
	q.array[q.rear] = val
	return
}

/*
	1、队列是一个有序可变， 可以用数组或链表实现FIFO
*/
func main() {
	q := &queue{
		maxSize: 4,
		array:   [4]int{},
		front:   -1,
		rear:    -1,
	}
	q.AddQueue(1)
	fmt.Println(q)
}
