package main

import (
	"fmt"
	"math/rand"
	_ "strconv"
	"sync"
	"time"
	"unsafe"
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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	cur1, cur2 := l1, l2
	res := &ListNode{}
	cur := res
	for cur1 != nil || cur2 != nil {
		if cur1 == nil {
			cur1 = &ListNode{Val: 0}
		}

		if cur2 == nil {
			cur2 = &ListNode{Val: 0}
		}

		total := cur1.Val + cur2.Val
		add := total / 10
		v := total % 10
		if add > 0 {
			if cur1.Next == nil {
				cur1.Next = &ListNode{
					Val: 0,
				}
			}
			cur1.Next.Val = cur1.Next.Val + add
		}

		fmt.Println(total, add, v)

		node := &ListNode{
			Val:v,
		}

		cur.Next = node
		cur = cur.Next

		cur1, cur2 = cur1.Next, cur2.Next
	}

	return cur.Next
}



func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	arr1, arr2 := nodeToArr(l1), nodeToArr(l2)
	len1, len2 := len(arr1), len(arr2)
	len := len1
	if len2 > len1 {
		len = len2
		diff := len2 - len1
		temarr := make([]int, diff, len2)
		arr1 = append(temarr, arr1...)
	} else if len1 > len2 {
		diff := len1 - len2
		temarr := make([]int, diff, len1)
		arr2 = append(temarr, arr2...)
	}

	res := &ListNode{}
	cur := res

	add := 0
	for i := len - 1; i >= 0; i-- {
		a1, a2 := arr1[i], arr2[i]
		sum := a1 + a2 + add
		add = sum / 10
		left := sum % 10
		fmt.Println(i, a1, a2, add, left)

		cur = &ListNode{
			Val: left,
			Next: cur,
		}

		fmt.Println(cur)
	}

	return cur
}

func nodeToArr(l *ListNode) []int {
	arr := []int{}
	cur := l
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}

	return arr
}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	xmap := make(map[int]int, len(nums1) * len(nums1))
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			idx := v1 + v2
			xmap[idx]++
		}
	}

	ans := 0
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			idx := 0 - v3 - v4
			if count, ok := xmap[idx]; ok {
				ans += count
			}
		}
	}

	return ans
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	mid, end := head, head
	for end.Next != nil && end.Next.Next != nil {
		mid = mid.Next
		end = end.Next.Next
	}

	end = mid.Next
	mid.Next = nil
	head = mid

	fmt.Println(mid, end)

	var tmp *ListNode
	for end != nil {
		t := end.Next
		end.Next = tmp
		tmp = end
		end = t
	}

	for head != nil && head.Next != nil {
		if head != tmp {
			return false
		}
		head = head.Next
		tmp = tmp.Next
	}

	return true
}

func re(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var res *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = res
		res, cur = cur, tmp
	}

	return res
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums) - 1
	if nums[start] > target && nums[end] < target {
		return -1
	}

	// {4,5,6,7,0,1,2} 0
	for start < len(nums) && end > 0 && start != end {
		if nums[start] == target {
			fmt.Println("start", len(nums), start)
			return start
		}

		if nums[end] == target {
			fmt.Println("end", len(nums), end)
			return end
		}

		start++
		end--
	}


	if nums[start] == target {
		return start
	}

	return -1
}

type apple struct {
	num int
	t []int
}

type egg struct {
	num int
	t []int
}

func (apple *apple) change(egg *egg)  {
	apple.num, egg.num = egg.num, apple.num
	apple.t, egg.t = egg.t, apple.t
}

// 洗牌算法
func shuffle(nums []int) []int  {
	rand.Seed(time.Now().UnixNano())
	for i := len(nums) - 1; i > 0; i-- {
		idx := rand.Intn(i)
		nums[idx], nums[i] = nums[i], nums[idx]
	}

	return nums
}

func selectByRand(nums []int, i int)[]int  {
	rand.Seed(time.Now().UnixNano())
	res := make([]int, 0, i)
	count := 0
	for ; i > 0 ; i-- {
		count++
		idx := rand.Intn(count) + 1
		if idx == count {
			res = append(res, nums[idx])
		}
	}

	return res
}

func monkey(sum int, dist int, carry int) int {
	if carry < 2 {
		return 0
	}

	if carry == 2 {
		if dist > 0 {
			return 0
		}
		if sum < 2 {
			return 0
		}

		return 1
	}

	if sum < dist {
		return 0
	}

	if carry > sum {
		 return sum - dist
	}

	for dist > 0 {
		if sum % carry <= 2 {
			sum = sum - (sum % carry)
		} else {
			sum -= (2* (sum -1) / carry) + 1
			dist--
		}
	}

	return sum
}

type mutex struct {

}

func main()  {
	time.AfterFunc(time.Millisecond * 100000, func() {
		fmt.Println("2 seconds ~")
	})
	fmt.Println(time.Now().Add(time.Millisecond * 100000).Unix(), time.Now().Unix())
	time.Sleep(time.Minute)
	// resp.activityDiscount += resp.originPrice - resp.originPrice*int32(mgr.activity.DiscountValue)/100
	// fmt.Println(172151 + 261178 + 23918 + 92652)

	// fmt.Println(monkey(50, 25, 25))
	//
	// fmt.Println(selectByRand([]int{1,2,3,4,5,6,7}, 3))
	//
	return

	a := &apple{
		num: 100,
		t: []int{1,2},
	}

	e := &egg{
		num: 0,
	}

	a.change(e)
	fmt.Println("apple:", a, "egg:", e)
	return


	fmt.Println(search([]int{3}, 3))
	// fmt.Println(search([]int{1,3,5}, 3))
	// fmt.Println(search([]int{4,5,6,7,0,1,2}, 6))

	return
	// node3 := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 		},
	// 	},
	// }
	//
	// x := re(node3)
	// fmt.Println(x)
	//
	// return
	//
	// fmt.Println(isPalindrome(node3))
	//
	// return

	nums1 := []int{1,2,3, 4, 5}
	nums2 := []int{-2,-1, -3}

	copy(nums1, nums2)

	fmt.Println(len(nums1), cap(nums1), nums1)

	return
	// nums3 := []int{-1,2, 0}
	// nums4 := []int{0,2, 0}
	//
	// fmt.Println(fourSumCount(nums1, nums2, nums3, nums4))
	// return
	// node1 := &ListNode{
	// 	Val: 7,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 4,
	// 			Next: &ListNode{
	// 				Val: 3,
	// 			},
	// 		},
	// 	},
	// }
	//
	// node2 := &ListNode{
	// 	Val: 5,
	// 	Next: &ListNode{
	// 		Val: 6,
	// 		Next: &ListNode{
	// 			Val: 4,
	// 		},
	// 	},
	// }

	// addTwoNumbers(node1, node2)
	//
	// return
	// s := "{}"
	// var xxxx ListNode
	// err := json.Unmarshal([]byte(s), &xxxx)
	// fmt.Println(xxxx, err)
	// fmt.Println(len(strings.Split(s, ",")))
	// return
	// x := []int{0,3,7,2,5,8,4,6,0,1}
	// fmt.Printf("%v, %p", x, x)
	// t(x)
	// fmt.Printf("%v, %p", x, x)
	// return
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