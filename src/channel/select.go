package main

import (
	"fmt"
	"time"
)


type ListNode struct {
	     Val int
	     Next *ListNode
	 }

func main()  {}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
	/*
		类似switch， 但是select语句是随机执行一个可运行的case
		如果无可执行case，要看是否有default， 如果没有，则进入阻塞， 直到有case可执行
	 */
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch2<- 10
	}()
	go func() {
		ch1 <- 2
	}()
	select {
	case num1 := <-ch1:
		fmt.Println("ch1:", num1)
	case num2 := <-ch2:
		fmt.Println("ch2:", num2)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout")
		default:
		fmt.Println("default")
	}

}
