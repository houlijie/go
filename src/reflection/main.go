package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

/**
	注意事项
	1、reflect.ValueOf.Kind 等同于reflect.IndexOf.Kind是一个常量
	2、
**/
func reflecttest(i interface{}) {
	rv := reflect.ValueOf(i)
	rt := reflect.TypeOf(i)
	fmt.Printf("rv=%v, rt=%v\n", rv, rt.Kind()) //rv={haha}, rt=main.stu
	var n1 int64 = 10
	// n2 := rv + n1 ////invalid operation: rv + n1 (mismatched types reflect.Value and int)
	n2 := rv.Int() + n1
	fmt.Println("n2=", n2)
}

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	numsMap := make(map[int]int, len(nums))
	for _, v := range nums {
		numsMap[v] = v
	}

	longist := 0
	for v := range nums {
		if _, ok := numsMap[v - 1]; !ok {
			currentNum := v
			currentLen := 1
			fmt.Println(numsMap[currentNum+1])
			for numsMap[currentNum+1] > 0 {
				currentNum++
				fmt.Println('a' , currentLen)
				currentLen++
				fmt.Println("bb", currentLen)
			}
			fmt.Println("yyyy", numsMap[currentNum+1], currentLen, longist)
			if longist < currentLen  {
				longist = currentLen
			}
		}
	}

	return longist
}

type ListNode struct {
	     Val int
	     Next *ListNode
}
/**
 * Definition for singly-linked list.
 *
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var mid, tail *ListNode
	pre := &ListNode{Next: head}
	pcur, cur, mcur := pre, head, mid
	for i:= 1; cur != nil; i++{
		tmp := cur.Next
		if i < left {
			pcur = pcur.Next
		}
		if i == left {
			mcur = cur
		}
		if i > right {
			tail = cur
			break
		}
		if i >= left {
			cur.Next = mid
			mid = cur
		}

		cur = tmp
	}

	mcur.Next = tail
	pcur.Next = mid

	return pre.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{}
	cur := newHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next=l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	fo(0, newHead)
	return newHead.Next
}

func fo(i int, head *ListNode)  {
	cur := head
	for cur != nil {
		fmt.Println("i=",i,  cur)
		cur = cur.Next
	}
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	len := 0
	cur := head
	for cur != nil {
		len++
		cur = cur.Next
	}

	fmt.Println(len, len / 2 + 1, head)
	for mid := len / 2 + 1; mid > 0; mid-- {
		head = head.Next
		fmt.Println(head)
	}

	return head
}

type noCopy struct {

}

func (*noCopy) Lock()  {
}
func (*noCopy) Unlock()  {

}
type nc struct {
	noCopy noCopy
	name string
}

var wg1 sync.WaitGroup

func odd(n int, ch chan int)  {
	for i := 0; i < n; i++ {
		<-ch
		if i %2 == 0 {
			fmt.Println("odd=", i)
		}
	}
	defer wg1.Done()
}

func even(n int, ch chan int)  {
	for i := 0; i < n; i++ {
		ch <- 1
		if i %2 == 1 {
			fmt.Println("even=" , i)
		}
	}
	defer wg1.Done()
}

func chick(cch chan struct{}, duch chan struct{}) {
	for i := 0; i < 10; i++ {
		<-cch
		duch <- struct{}{}
		fmt.Println("chick")
	}
	wg1.Done()
}

func duck(duch chan struct{}, doch chan struct{})  {
	for i := 0; i < 10; i++ {
		<-duch
		doch <- struct{}{}
		fmt.Println("duck")
	}
	wg1.Done()
}

func dog(doch, cch chan struct{})  {
	for i := 0; i < 10; i++ {
		cch <- struct {}{}
		<-doch
		fmt.Println("dog")
	}
	wg1.Done()
}



// 遍历打印鸡鸭狗 10遍
func p()  {
	wg1.Add(3)
	duch := make(chan struct{})
	doch := make(chan struct{})
	cch := make(chan struct{})

	go chick(cch, duch)
	go duck(duch, doch)
	go dog(doch, cch)
	wg1.Wait()
}

// 10个协程
func ten()  {
	ch := make(chan int, 10)
	wg1.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				select {
				case i, ok := <-ch:
					fmt.Println(i, ok)
					if !ok {
						break
					}
				}
			}
			wg1.Done()
		}()
	}
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	// close(ch)

	wg1.Wait()
}

func main() {
	go func() {
		panic("hahahaha")
	}()
	go func() {
		panic("hihi")
	}()

	time.Sleep(time.Second * 5)
	return
	// ch := make(chan int)
	// wg1.Add(2)
	// go even(10, ch)
	// go odd(10, ch)
	// wg1.Wait()
	//
	// return

	// x := make(map[int]int, 10)
	// for i := 0; i < 10; i++ {
	// 	x[i] = i
	// }
	// fmt.Println("x1", len(x))
	// for i, _ := range x {
	// 	if i % 2 == 0 {
	// 		delete(x, i)
	// 		x[10] = 1
	// 	}
	// }
	//
	// fmt.Println("x2", len(x), x)
	//
	// for i := range x{
	// 	delete(x, i)
	// }
	//
	// fmt.Println("x2", len(x))
	//
	// return
	// ctx := context.Background()
	// ctx1, cancel:= context.WithDeadline(ctx, time.Now().Add(time.Second * 13))
	// defer  cancel()
	//
	// x, ok := ctx1.Deadline()
	// fmt.Println(x, ok)
	//
	//
	//
	//
	// ctx2 := context.WithValue(ctx1, "name", "hlj")
	// ctx3 := context.WithValue(ctx2, "age", "1")
	// context.WithDeadline(ctx1, time.Now().Add(time.Second * 2))
	//
	//
	// fmt.Println(ctx1.Value("name"), ctx2.Value("name"), ctx3.Value("name"))
	//
	// ch := make(chan int, 5)
	// for i := 0; i < 5; i++ {
	// 	ch <- i
	// }
	//
	// var wg sync.WaitGroup
	// wg.Add(2)
	//
	// go func() {
	// 	time.Sleep(time.Microsecond)
	// 	close(ch)
	// 	wg.Done()
	// }()
	//
	// go func() {
	// 	for v := range ch {
	// 		fmt.Println("v=xxyy", v)
	// 	}
	// 	wg.Done()
	// }()
	//
	// wg.Wait()
	fmt.Println("done")
	return

	sq := "abcdefg"
	sq2 := "后理解s"
	sqr := sq2//[]rune(sq2)
	// for i, i2 := range sqr {
	// }
	fmt.Println(len(sqr), fmt.Sprintf("%c***%c", sqr[0], sqr[len(sqr) - 1]))

	// fmt.Println(len(sq2), len(sq))
	s, err := strconv.ParseBool(sq)
	fmt.Println(s, err)
	return

	l1 := &ListNode{
		1,&ListNode{
			2, &ListNode{
				4, nil,
			},
		},
	}

	fmt.Println(middleNode(l1))
	return
	l2 := &ListNode{
		1,&ListNode{
			3, &ListNode{
				4, nil,
				},

		},
	}
	fmt.Println(mergeTwoLists(l1, l2))
	return
	l := &ListNode{
		1,&ListNode{
			2, &ListNode{
				3, &ListNode{4, &ListNode{
					5, nil},
				},
			},
		},
	}
	fmt.Println(reverseBetween(l, 2, 4))

	return
	a := "abc"
	fmt.Println(a[2:3],  len(a))
	return
	// a := longestConsecutive([]int{1,2,3})
	errors.New()

	// fmt.Println(a)
	return
	// var s2 int64 = 10
	// s := reflect.ValueOf(s2)
	// fmt.Println(s.Type())
	// fmt.Println("s2=", s2)
	// reflecttest(s2)
}
