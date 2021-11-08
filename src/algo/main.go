package main

import (
	"fmt"
	_ "strconv"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func compareVersion(version1 string, version2 string) int {
	v1, v2 := 0, 0

	for _, v := range version1 {
		if v != '.' {
			v1 = v1 * 10 + int(v - '0')
		}
	}

	for _, v := range version2 {
		if v != '.' {
			v2 = v2 * 10 + int(v - '0')
		}
	}

	fmt.Println(v1, v2)

	if v1 > v2 {
		return 1
	}

	if v1 < v2 {
		return -1
	}

	return 0
}

func t(x []int)  {
	x[0] = 2
	// x = append(x, 2)
	fmt.Printf("%v, %p", x, x)
}

func main()  {
	x := []int{0,3,7,2,5,8,4,6,0,1}
	fmt.Printf("%v, %p", x, x)
	t(x)
	fmt.Printf("%v, %p", x, x)
	return
	// fmt.Println(compareVersion("1.01", "1.001"))
	// return
	// head := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 		},
	// 	},
	// }
	// cur := head
	// count := 0
	// var res *ListNode
	// rand2.Seed(time.Now().UnixNano())
	// for cur != nil {
	// 	count++
	// 	rand := rand2.Intn(count) + 1
	// 	fmt.Println(rand, count)
	// 	if rand == count {
	// 		res = cur
	// 	}
	// 	cur = cur.Next
	// }
	//
	// fmt.Println(res)
	// return
	//
	//
	// a, b := 2.556, 0.0
	//
	// c := int64((a+b) * 100.0)
	//
	// y := (a+b) * 100.0
	//
	// t, x := math.Modf(a)
	// fmt.Println(t, x)
	//
	//
	// fmt.Println(c, y, "ceil:", math.Ceil(y), "floor:", math.Floor(y))


}