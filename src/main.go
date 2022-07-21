package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func t() {
	a := []int{-1, -2, -3, 0, 6, 7, 8, 1}
	max := 0
	s, e := 0, 1
	sum := 0
	for s < e && e < len(a) {
		sum += a[s]
		if sum == 0 {
			len := e - s
			if max > len {
				max = len
			}
		}
	}
}

func t1() {
	fmt.Println("this is t1")
}

func t2() {
	fmt.Println("this is t2")
	panic("this is t2 panic")
}

type Lock struct {
	ch chan struct{}
}

func NewLock() *Lock {
	ch := make(chan struct{}, 1)
	return &Lock{
		ch: ch,
	}
}

func t3(a []int) []int {
	l, r := 0, 0
	for r < len(a) {
		if a[r] != 0 {
			a[l], a[r] = a[r], a[l]
			l++
		}
		r++
	}
	return a
}

func t4(s string) string {
	res := make([]string, 0, len(s))
	start, end := 0, 0
	for end < len(s) {
		if s[end] == ' ' {
			if end > start {
				res = append(res, s[start:end])
			}
			end++
			start = end
			continue
		}
		end++
	}

	if end > start {
		res = append(res, s[start:end])
	}

	resLens := len(res) - 1

	var ret strings.Builder
	ret.Grow(len(s))
	ret.WriteString(res[resLens])
	for i := resLens - 1; i >= 0; i-- {
		ret.WriteString(" ")
		ret.WriteString(res[i])
	}

	return ret.String()
}

func t5(q string, replicas []int) int {

	if true {
		fmt.Println("ssss")
		return 1
	}
	defer func() {
		fmt.Println("defer...")
	}()
	c := make(chan int)
	callback := func(i int) {
		c <- replicas[i]
	}

	for i := range replicas {
		go callback(i)
	}

	return <-c
}

type name struct {
	new func(i int) interface{}
}

func threeSum1(nums []int) [][]int {
	fmt.Println(nums)
	sort.Ints(nums)
	res := make([][]int, 0, 0)
	l := len(nums)
	for i := 0; i < l; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		thirdIndex := l - 1
		target := 0 - nums[i]
		for j := i + 1; j < l; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < thirdIndex && nums[thirdIndex] > target-nums[j] {
				thirdIndex--
			}

			if j == thirdIndex {
				break
			}

			if target-nums[j] == nums[thirdIndex] {
				res = append(res, []int{nums[i], nums[j], nums[thirdIndex]})
			}
		}
	}

	return res
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		third := n - 1
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for third > j && 0-nums[i]-nums[j] < nums[third] {
				third--
			}
			if third == j {
				break
			}
			if nums[i]+nums[j]+nums[third] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[third]})
			}
		}
	}

	return ans
}

func rand50() int {
	rand.Seed(time.Now().UnixNano())

	return (rand.Intn(9) + 1) & 1
}

func rand75() int {
	return rand50() | rand50()
}

var (
	x = 1
	b = x + 1
)

func hello1(num ...int) {
	num[0] = 19
	num = append(num, 10)
}

type xx struct {
	age  int
	name string
	xxx
}

type xxx struct {
	sss int
}

func mt1(m map[int]xx) {
	for i, v := range m {
		v.name = "lisi"
		m[i] = v
	}
	for i := 0; i < 20; i++ {
		m[i] = xx{
			age: 23,
		}
	}

	fmt.Println(m[1].age)
}

func g1() {
	fmt.Println("g1")
	panic("g1 panic")
}

type inte int

func (i *inte) Add(b inte) inte {
	ret := *i + b
	return ret
}

var mu sync.RWMutex

type iti struct {
	name string
	age  int
}
type it struct {
	iti
	gender int
}

func it1(i interface{}) {
	switch ret := i.(type) {
	case *iti:
		fmt.Println("ret", ret)
	case *int:
		fmt.Println("ret it ", ret)
	}
}

func f1() int {
	time.Sleep(time.Second * 1)
	return 1
}

func f2() int {
	time.Sleep(time.Second * 2)
	return 1
}

func f3() int {
	time.Sleep(time.Second * 3)
	return 1
}

func f4(a int, b int, c int) {
	return
}

func sh(nums []int) []int {
	start, end := 0, 0
	for end < len(nums) {
		if nums[end] != 0 {
			nums[start], nums[end] = nums[end], nums[start]
			start++
		}
		end++
	}

	return nums
}

func reversePrefix(word string, ch byte) string {
	index := -1
	wordByte := []byte(word)
	for i, v := range wordByte {
		if v == ch {
			index = i
			break
		}
	}

	for i := 0; i < index; i++ {
		wordByte[i], wordByte[index] = wordByte[index], wordByte[i]
		index--
	}

	return string(wordByte)
}

func main() {
	fmt.Println(reversePrefix("abcdesf", 'd'))
	fmt.Println(sh([]int{1, 2, 0, 0, -1, 0, 4}))
	return
	time1 := time.Now()
	f4(f1(), f2(), f3())
	time2 := time.Now().Sub(time1).Seconds()
	fmt.Println(time2)
	return
	mx := make(map[string]int, 3)
	mx["123"] = 1
	mx["456"] = 2
	mx["789"] = 3
	mx2 := make(map[int]string, 3)
	for s, i := range mx {
		mx2[i] = s
	}
	fmt.Println(mx2)
	return
	t := &it{
		iti: iti{
			name: "1",
			age:  0,
		},
		gender: 0,
	}
	it1(t)
	return
}
func A1() {
	mu.RLock()
	defer mu.RUnlock()
	B2()
}
func B2() {
	time.Sleep(3 * time.Second)
	C()
}
func C() {
	mu.RLock()
	defer mu.RUnlock()
}

// func main() {
// 	var b
// 	context.Background()
// 	b = true
// 	b = 1 == 2
//
// 	return
// 	defer
//
// 	func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	go g1()
// 	time.Sleep(time.Second * 2)
// 	return
// 	// map1 := make(map[int]xx, 1)
// 	// map1[1] = xx{
// 	// 	age:  1,
// 	// 	name: "zhangsan",
// 	// }
// 	// mt1(map1)
// 	// fmt.Println(map1)
// 	// return
// 	// al1 := [3]int{1, 2, 3}
// 	// fmt.Println(cap(&al1))
// 	// return
// 	sl1 := []int{1, 2, 3}
// 	sl2 := sl1[:2]
// 	hello1(2, 3, 4, 5)
// 	fmt.Println(sl1, cap(sl1), sl2, cap(sl2), len(sl2))
// 	// sn1 := struct {
// 	// 	age  int
// 	// 	name string
// 	// }{
// 	// 	age:  11,
// 	// 	name: "lisi",
// 	// }
// 	// sn2 := struct {
// 	// 	age  int
// 	// 	name string
// 	// }{
// 	// 	age:  11,
// 	// 	name: "lisi",
// 	// }
// 	//
// 	// m1 := []int{1}
// 	// sm1 := struct {
// 	// 	age int
// 	// 	m   []int
// 	// }{
// 	// 	age: 1,
// 	// 	m:   m1,
// 	// }
// 	// sm2 := struct {
// 	// 	age int
// 	// 	m   []int
// 	// }{
// 	// 	age: 1,
// 	// 	m:   m1,
// 	// }
// 	// if sn1 == sn2 {
// 	// 	fmt.Println("======")
// 	// }
// 	// if sm1 == sm2 {
// 	// 	fmt.Println("mmmm")
// 	// }
// 	// for i := 0; i < 50; i++ {
// 	// 	fmt.Println(rand75())
// 	// }
// 	// fmt.Println(rand.Intn(9) + 1
//
// 	// for
// 	return
// 	fmt.Println(threeSum([]int{-1, -1, 0, 1, 2, -1, -4}))
// 	// fmt.Println(threeSum([]int{0, 0, 0, 0}))
// 	return
//
// 	// i := 0
// 	// for {
// 	//
// 	// 	go func(i int) {
// 	// 		url := fmt.Sprintf("https://www.xxx%d.com/", i)
// 	// 		resp, err := http.Get(url)
// 	// 		fmt.Println(resp, err)
// 	// 		if err != nil {
// 	// 			fmt.Printf("http.Get err: %v\n", err)
// 	// 		}
// 	// 	}(i)
// 	// 	i++
// 	// 	time.Sleep(time.Second * 1)
// 	// 	fmt.Println("goroutines: ", runtime.NumGoroutine())
// 	// }
// 	s := " hello world "
// 	// for i := 0; i < 3; i++ {
// 	t5(s, []int{0, 1, 2, 3, 4})
// 	// 	fmt.Println(runtime.NumGoroutine())
// 	// }
// 	// time.Sleep(time.Second * 3)
// 	// return
// 	// fmt.Println(t4(s))
// 	return
// 	a := make([]int, 0, 8)
// 	a = append(a, []int{1, 2, 0, 0, -1, 0, 4}...)
// 	fmt.Println(t3(a))
// 	return
// 	// sync.Mutex{}
// 	defer t2()
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("this is t1")
// 	}()
//
// 	go func() {
// 		defer wg.Done()
// 		defer func() {
// 			if err := recover(); err != nil {
// 				fmt.Println("t2 panic")
// 			}
// 		}()
// 		fmt.Println("this is t2")
// 		panic("this is t2 panic")
// 	}()
// 	wg.Wait()
// 	time.Sleep(time.Second * 2)
// 	return
// 	// tt := new(map[uint64]uint64)
// 	// if tt == nil {
// 	// 	fmt.Println("nil")
// 	// }
// 	// // tt[1] = 1
// 	//
// 	// fmt.Println("not nil", len(*tt))
// 	//
// 	// return
// 	// a := uint64(12)
// 	// s := 9
// 	// x := atomic.AddUint64(&a, uint64(-s))
// 	// fmt.Println(a, x)
// 	// return
// 	// var wg sync.WaitGroup
// 	// var o sync.Once
// 	// wg.Add(3)
// 	// for i := 0; i < 3; i++ {
// 	// 	go func() {
// 	// 		defer wg.Done()
// 	// 		o.Do(t1)
// 	// 	}()
// 	// }
// 	// wg.Wait()
// 	//
// 	// return
// 	//
// 	// wg1 := wg
// 	// wg1.Add(1)
// 	// wg1.Done()
// 	// wg1.Wait()
// 	//
// 	// wg.Add(2)
// 	// go func() {
// 	// 	t(time.Minute * 2)
// 	// 	wg.Done()
// 	// }()
// 	// go func() {
// 	// 	t(time.Minute * 3)
// 	// 	wg.Done()
// 	// }()
// 	//
// 	// wg.Wait()
// }

//
// func t1() {
// 	select {}
// }
// func t(t time.Duration) {
// 	db, err := sql.Open("mysql", "root:root@(localhost)/test")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer db.Close()
// 	if err := db.Ping(); err != nil {
// 		fmt.Println(err)
// 	}
//
// 	tx, err := db.Begin()
// 	if err != nil {
// 		fmt.Println("begin", err)
// 		return
// 	}
// 	rows, err := db.Exec("update t set a= 1 where id = 1 ")
// 	// rows, err := db.Query("select * from t where id = 1 ")
// 	fmt.Println("===", rows, err)
//
// 	// time.Sleep(t)
// 	tx.Commit()
// }
