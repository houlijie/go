package main

import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

func Go(f func())  {
	go func() {
		defer func() {
			if  err := recover();err != nil {
				fmt.Println("go panic")
			}
		}()
		f()
	}()
}

func moveZeroes(nums []int)  {
	for i, j := 0, 1; j < len(nums); j++ {
		fmt.Println(i, j, nums)
		if nums[i] != 0 {
			i++
			continue
		}
		if nums[j] != 0 {
			fmt.Println("====", i, j, nums)
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	fmt.Println(nums)
}

func twoSum(nums []int, target int) []int {
	for i:= 0; i<len(nums)-1; i++ {
		for j := 1; j < len(nums); j++ {
			fmt.Println(i, j)
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func isValid(s string) bool {
	l := len(s)
	if l % 2 != 0 {
		return false
	}

	m := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}

	stack := make([]rune, 0, l/2)
	for _, v := range s {
		if pair, ok := m[v]; !ok {
			stack = append(stack, v)
		} else {
			fmt.Println(stack)
			slen := len(stack)
			if slen == 0 {
				return false
			}

			last := stack[slen -1]
			if last == pair {
				stack = stack[:slen-1]
			}
		}
	}

	return len(stack) == 0
}

type empty struct {}

type name struct {
	a int32
	// b int32
	c []int32
	empty
}

func test() (resp *name, err error) {
	resp.a = 10

	return resp, nil
}


func A(i int) (ret int)  {
	defer func(i int) {
		ret++
		fmt.Println(i)
	}(i)
	ret +=i
	return
	
}

func  main() {
	c1 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer  wg.Done()
		for  {
			select {
			case v, ok := <-c1:
				fmt.Println("g1 closed", v, ok)
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case v, ok:= <-c1:
				fmt.Println("g2 closed", v, ok)
				return
			case <-time.After(time.Microsecond):
				fmt.Println("timeout")
				return
			}
		}
	}()

	go func() {
		c1 <- 1
		// close(c1)
	}()



	wg.Wait()

	// fmt.Println(A(1))
	return
	ch := make(chan int, 1)

	go func() {
		for {
			fmt.Println(111)
			select {
			case <-ch:
				fmt.Println("xxxxxxx")
				continue
				fmt.Println("xxxxxxx")
			}
		}
	}()
	ch <-1

	time.Sleep(time.Second * 13)
	return

	go func() {
		for {
			fmt.Println("y==")
			select {
			case <-ch:
				fmt.Println("yyy")
				return
			default:
				fmt.Println("y default")
			}
		}
	}()
	time.Sleep(time.Second * 1)
	close(ch)

	time.Sleep(time.Second * 1)
	return
	ss := []uint64{1, 2, 3, 4, 5, 6}
	//
	//fmt.Println(len(ss), cap(ss))
	sss := ss[1:2:4]

	fmt.Println(sss, len(sss), cap(sss))

	sss = ss[1:2:3]

	fmt.Println(sss, len(sss), cap(sss))
	return
	xxx, _ := test()
	fmt.Println("====", xxx)
	// xxx.c = []int32{1, 2, 3}
	return
	var name name
	fmt.Println("===", unsafe.Sizeof(name), unsafe.Alignof(name))
	return

	// sss := "abc=111&mall_id=-123&acc=233"
	// q, _:= url.ParseQuery(sss)
	// x, err := strconv.ParseUint(q.Get("mall_id"), 10, 64)
	// fmt.Print(q, "====", x, err)
	// return
	a := []int{1, 2, 3}
	b := a[1:]
	b = append(b, 4)
	fmt.Println(b, len(b), cap(b))
	return

	fmt.Println(isValid("([]){}"))
	return
	moveZeroes([]int{1,0,0,0,0,1})
	return
	// s := "" + "123"
	// ss, err := strconv.Atoi(s)
	// if err != nil || ss < 0 {
	// 	fmt.Println(ss, err, "eee")
	// }
	//
	// s1 := strings.TrimRight(s, " ")
	// l := 0
	// for i := len(s1) -1; i > 0; i-- {
	// 	if s1[i] == ' ' {
	// 		break
	// 	}
	// 	l++
	// }
	//
	// fmt.Println(l)
	// return
	// Go(func() {
	// 	panic("gooooo")
	// })
	// go func() {
	// 	panic("hahahaha")
	// }()
	// go func() {
	// 	panic("hihi")
	// }()

	time.Sleep(time.Second * 5)
	// var wg sync.WaitGroup
	// wg.Add(2)
	// ch := make(chan struct{})
	// go func() {
	// 	defer wg.Done()
	// 	for i := 1; i <= 100 ; i++ {
	// 		ch <- struct{}{}
	// 		if  i % 2 == 1 {
	// 			fmt.Println(i)
	// 		}
	//
	// 	}
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	for i := 1; i <= 100 ; i++ {
	// 		<- ch
	// 		if i % 2 == 0 {
	// 			fmt.Println(i)
	// 		}
	// 	}
	// }()
	// wg.Wait()
}
