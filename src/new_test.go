package main

import (
	"fmt"
	"sync"
	"unsafe"

	"github.com/gin-gonic/gin"
)

type tester interface {
	test()
	haha() string
	string() string
}
type data struct{}

func (*data) test() {}
func (data) string() string {
	return "aaa"
}
func (data) haha() string {
	return "haha"
}

func hello(num ...int) {
	num[0] = 18
	fmt.Printf("%p, %+v", num, num)
}

const i = 100

var j = 123

type ae struct {
}
type s1 struct {
	a int32
	c int64
	b int8
	ae
}

// type Context interface {
// 	// done返回一个channel， 当context 被取消/超时，chanle已关闭
// 	Done() <-chan struct{}
// 	// err表示当done channel被关闭时， 这个context为什么被取消
// 	Err() error
//
// 	// 会返回一个超时时间， runtime获取超时时间后， 可以对某些io操作设定超时时间
// 	DeadLine() (deadline time.Time, ok bool)
//
// 	// value可以让rotine共享一些数据
// 	Value(key interface{}) interface{}
// }

func incr(d int) (ret int) {
	// defer fmt.Println("====", ret)
	// // defer func(ret int) {
	// // 	ret++
	// // }(ret)
	defer func() {
		ret++
		fmt.Println("func==", ret)
	}()
	return d

	// t := 5
	// defer func() {
	// 	t += 5
	// }()
	//
	// return t
	// defer func(r int) {
	// 	r += 5
	// }(ret)
	//
	// return 1

}

// 交替打印100
func p1() {
	ch := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 1; i <= 100; i++ {
			<-ch
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
		wg.Done()
	}()

	go func() {
		for i := 1; i <= 100; i++ {
			ch <- struct{}{}
			if i%2 == 1 {
				fmt.Println(i)
			}
		}
		wg.Done()
	}()

	wg.Wait()
}

func p() {

	var wg sync.WaitGroup
	ch1 := make(chan struct{})
	ch2 := make(chan struct{}, 1)
	ch3 := make(chan struct{})

	wg.Add(3)
	go func() {
		defer wg.Done()
		do("dog", ch3, ch1)

	}()

	go func() {
		defer wg.Done()
		do("cat", ch2, ch3)
	}()

	go func() {
		defer wg.Done()
		do("fish", ch1, ch2)
	}()
	ch2 <- struct{}{}
	wg.Wait()
}

func do(s string, in chan struct{}, out chan struct{}) {
	for i := 0; i < 100; i++ {
		<-out
		in <- struct{}{}
		fmt.Println(s)
	}
}

type a []int

func t1(a *[]int) {
	*a = append(*a, 1, 2, 3, 4)
	return
}

func t2(a []int) {
	a = append(a, 1, 2, 3, 4)
	return
}

func p2(i int) int {
	defer func(i int) {
		i++
	}(i)

	// return i++
	return i
}

func p3(i int) int {
	defer func() {
		i++
	}()

	// return i++
	return i
}

var wg sync.WaitGroup

func f1() int {
	ret := 5
	defer func(ret int) {
		ret += 10
	}(ret)

	return ret
}
func f2() (ret int) {
	ret = 5
	defer func(ret int) {
		ret += 5
	}(ret)
	return ret
}

func f3() (ret int) {
	ret = 5
	defer func() {
		ret += 10
	}()
	return ret
}

func A() {
	defer println("A")
	wg.Add(1)
	go B()
	wg.Wait()
}

func B() {
	defer println("B")
	panic("this is b")
	wg.Done()
}

func D() {
	panic("xxx")
	fmt.Println("dddd")
}

func C() {
	defer D()
	panic("this is c")
}

func m1(g *gin.Engine) {
	fmt.Println("m1 start")
}

func M2(x map[int]int) {
	x[2] = 3
	x[3] = 3
}

func M3(x []int) {
	x = append(x, 3)
}

func main() {
	m2 := map[int]int{
		1: 1,
		2: 2,
		3: 2,
		4: 2,
		5: 2,
	}

	for i2, i3 := range m2 {
		delete(m2, 3)
		fmt.Println(i2, i3)
	}
	// fmt.Printf("%#v", m2)
	//
	M2(m2)
	fmt.Println(m2)
	return

	m3 := make([]int, 0, 8)
	m3 = append(m3, []int{1, 2}...)

	M3(m3)
	p := unsafe.Pointer(&m3[1])
	p1 := uintptr(p) + 8
	fmt.Println(m3, cap(m3), *(*int)(unsafe.Pointer(p1)))
	return
	// wg.Add(50)
	// for i := 0; i < 50; i++ {
	// 	go func() {
	// 		defer wg.Done()
	// 		sx = append(sx, i)
	// 	}()
	// }
	// wg.Wait()
	// fmt.Println(sx, len(sx))
	// // fmt.Println(f1(), f2(), f3())
	// // A()
	// // C()
	// return
	// a := 5
	// switch a {
	// case 5:
	// 	fmt.Println(a)
	// case 3:
	// 	fmt.Println(a)
	// }
	// return
	var wg sync.WaitGroup
	wg.Add(6)
	ch := make(chan int, 100)
	stopCh := make(chan struct{})
	toStopCh := make(chan struct{}, 1)
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Println(i)
				select {
				case <-stopCh:
					return
				default:
					ch <- i
				}
			}
		}(i)
	}

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stopCh:
					return
				case v := <-ch:
					if v == 9 {
						fmt.Println("==", v)
						toStopCh <- struct{}{}
						return
					}
				default:

				}
			}
		}()
	}

	go func() {
		defer wg.Done()
		select {
		case <-toStopCh:
			close(stopCh)
		}
	}()

	wg.Wait()
	return
	// ch := make(chan int)
	// // close(ch)
	// go func() {
	// 	for v := range ch {
	// 		fmt.Println(v)
	// 	}
	// }()
	// return
	// s1 := make([]int, 0, 2000)
	// println(s1)
	// return
	// defer A()
	// return
	// i := 1
	// fmt.Println(p2(i), p3(i))
	// chq := make(chan int, 1)
	// close(chq)
	// fmt.Println(<-chq)
	// return
	// a := make([]int, 0, 4)
	// a = append(a, 1)
	// fmt.Println(a, cap(a))
	// // t1(&a)
	// t2(a)
	// fmt.Println("==", a, cap(a))
	//
	// go fmt.Println(p2(1))
	// return
	// ctx := context.Background()
	// ctx1 := context.WithValue(ctx, "hhahha", "hlj1")
	// ctx2 := context.WithValue(ctx1, "hhahha1", "hlj2")
	// fmt.Println(ctx2.Value("hhahha"))
	// return
	// m1 := make(map[int]int, 2)
	// go func() {
	// 	for {
	// 		m1[1] = 2
	// 	}
	// }()
	// go func() {
	// 	for {
	// 		_, _ = m1[1]
	// 	}
	// }()
	// fmt.Println("main===", incr(3))
	// return
	// s11 := s1{}
	// fmt.Println(unsafe.Sizeof(s11))
	// return
	// a := []int{1, 2, 3, 4, 5}
	// r := make([]int, 0)
	// for i, v := range a {
	// 	if i == 0 {
	// 		a = append(a, 6, 7)
	// 	}
	// 	r = append(r, v)
	// }
	// fmt.Println(r)
	// return
	// x := []int{5: 2, 3, 8: 1}
	// fmt.Println(x)
	// var ch chan int
	// ch = make(chan int, 1)
	// close(ch)
	// ch <- 0
	// <-ch
	// m := map[int]string{0: "zero", 1: "one"}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }
	// fmt.Println(&j, j)
	// a = []int{1, 2, 3}
	//
	// b := a[1:]
	// fmt.Println(b[1])
	//
	// return
	// i := -5
	// j := +5
	// fmt.Printf("%+d, %+d", i, j)
	// s := make(map[string]int)
	// delete(s, "h")
	// var i interface{}
	// if i == nil {
	// 	fmt.Println("interface is nil")
	// 	return
	// }
	// fmt.Println(s["h"], cap(i))
	// return
	// var d data
	//
	// // var t tester = d // 错误：test()不属于data的方法集，所以不支持转换
	// var t tester = &d // 结构类型是可以直接转换为接口类型
	// t.test()
	// println(t.string(), d.haha())
}

// func callA(ctx context.Context, now time.Time) {
// 	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
// 	defer cancel()
// 	select {
// 	case <-ctx.Done():
// 		fmt.Println("aaaaaa")
// 	}
//
// 	fmt.Println("aaa time use ", time.Now().Unix()-now.Unix())
// }
//
// func callB(ctx context.Context, now time.Time) {
// 	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
// 	defer cancel()
// 	select {
// 	case <-ctx.Done():
// 		fmt.Println("bbbbb")
// 	}
//
// 	fmt.Println("bbb time use ", time.Now().Unix()-now.Unix())
// }
//
// func hello(i *int) int {
// 	defer func() {
// 		*i = 19
// 	}()
// 	return *i
// }
//
// func returnValues() int {
// 	var result int
// 	defer func() {
// 		result++
// 		fmt.Println("defer", "res", result)
// 	}()
//
// 	panic("this is panic")
// 	return result
// }
//
// func namedReturnValues() (result int) {
// 	defer func() {
// 		result++
// 		fmt.Println("defer")
// 	}()
// 	return
// }
//
// type Foo struct {
// 	v int
// }
//
// type Bar struct {
// 	v int
// }
//
// func NewFoo(n *int, pre string) Foo {
// 	println("newfoo==", pre, *n)
// 	return Foo{}
// }
//
// func (Foo) Bar(n *int) Bar {
// 	println("Bar===", *n)
// 	return Bar{}
// }
//
// func (Bar) Tar(n *int) {
// 	println("Tar===", *n)
// }
//
// type T int
//
// func (t T) M(n int) T {
// 	print(n)
// 	return t
// }
//
// type Add func(int, int) int
//
// // 4， 2
// func combine(n int, k int) [][]int {
// 	ans := make([][]int, 0)
// 	path := make([]int, 0)
// 	return dfs(n, k, 1, path, ans)
// }
//
// func dfs(n, k, begin int, path []int, ans [][]int) [][]int {
// 	if len(path) == k {
// 		x := make([]int, len(path))
// 		copy(x, path)
// 		ans = append(ans, x)
// 		return ans
// 	}
//
// 	for i := begin; i <= n; i++ {
// 		path = append(path, i)
// 		ans = dfs(n, k, i+1, path, ans)
// 		path = path[:len(path)-1]
// 	}
//
// 	return ans
// }
//
// func main() {
// 	xxx := combine(4, 2)
// 	fmt.Println(xxx)
// 	return
// 	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
//
// 	return
// 	fmt.Println(make([]int, 1, 2))
// 	return
// 	for i := 0; i < 5; i++ {
// 		switch i {
// 		case 0:
// 			fmt.Println(i)
// 		case 1, 3:
// 			fallthrough
// 		case 2:
// 			fmt.Println(2)
// 		}
//
// 	}
// 	// var f Add
// 	// defer f(1, 2)
//
// 	// fmt.Println("end")
// 	// return
// 	// var t T
// 	// defer t.M(1).M(2)
// 	// t.M(3).M(4)
// 	// return
// 	// var x = 1
// 	// var p = &x
// 	// defer NewFoo(p, "defer")
// 	// x = 2
// 	// p = new(int)
// 	// NewFoo(p, "main")
// 	// return
// 	// eg, _ := errgroup.WithContext(context.Background())
// 	// for i := 1; i < 3; i++ {
// 	// 	i := i
// 	// 	eg.Go(func() error {
// 	// 		s := time.Now().Unix()
// 	// 		time.Sleep(time.Duration(i) * time.Second)
// 	// 		if i == 1 {
// 	// 			return errors.New("i = 1 err")
// 	// 		}
// 	//
// 	// 		fmt.Println("i=", i, ", time:", time.Now().Unix()-s)
// 	// 		return nil
// 	// 	})
// 	// }
// 	//
// 	// res := eg.Wait()
// 	// fmt.Println("res", res)
// 	//
// 	// return
// 	// s1 := returnValues()
// 	// fmt.Println(s1)
// 	// // s2 := namedReturnValues()
// 	// // return
// 	// i := 10
// 	// j := hello(&i)
// 	// fmt.Println(i, j)
// 	// return
// 	//
// 	// now := time.Now()
// 	// ctx, cal := context.WithTimeout(context.Background(), time.Second*5)
// 	// defer cal()
// 	//
// 	// var wg sync.WaitGroup
// 	// wg.Add(2)
// 	// go func() {
// 	// 	callA(ctx, now)
// 	// 	wg.Done()
// 	// }()
// 	//
// 	// go func() {
// 	// 	callB(ctx, now)
// 	// 	wg.Done()
// 	// }()
// 	//
// 	// wg.Wait()
// 	// fmt.Println("main time use ", time.Now().Unix()-now.Unix())
//
// }
//
// func lengthOfLongestSubstring(s string) int {
// 	strLen := len(s)
// 	if strLen < 2 {
// 		return strLen
// 	}
//
// 	m := make(map[byte]struct{}, strLen)
//
// 	maxLen := 0
// 	start, end := 0, 1
// 	for end < strLen {
// 		str := s[end]
// 		if _, ok := m[str]; ok {
// 			delete(m, str)
// 			fmt.Println("==sss==", start, end)
// 			maxLen = max(end-start, maxLen)
//
// 			start = start + 1
// 			end = start + 1
// 			fmt.Println("====", start, end, maxLen)
// 		}
//
// 		m[str] = struct{}{}
// 		end++
// 		continue
//
// 	}
//
// 	fmt.Println("===eee=", start, end, maxLen)
//
// 	maxLen = max(end-start, maxLen)
//
// 	return maxLen
// }
//
// func max(a, b int) int {
// 	if a >= b {
// 		return a
// 	}
//
// 	return b
// }
