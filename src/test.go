package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func assert(i interface{}) {

	switch i.(type) {
	case string:
		fmt.Println(i.(string))
	case int:
		fmt.Println(i.(int))
	case []int:
		fmt.Println(i.([]int))
	case map[int]int:
		fmt.Println(i.(map[int]int))
	default:
		fmt.Printf("格式错误: %T", i)
	}
}

const (
	// AA, BB = iota, iota
	// A      = iota
	// B
	// B      = "aaa"
	// C      = iota + 2
	// _
	// D
	// F
	E = iota
)

const (
	AA, BB = iota + 2, iota + 3
	A, B
)

const BizGroup int = 0xFFFF

var c = make(chan int)

// var a int
// func f() {
// 	a = 1
// 	<-c
// }

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("i===", i)
		chnl <- i
	}

	fmt.Println("close")
	close(chnl)
}

func a() {
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}

type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch t := v.(type) {
	case *student:
		a := v.(*student)
		fmt.Println(a.Name)
	default:
		fmt.Println(t)
	}
}

type Student struct {
	number   int
	realname string
	age      int
}

func (t *Student) String() string {
	return fmt.Sprintf("学号: %d\n真实姓名: %s\n年龄: %d\n", t.number, t.realname, t.age)
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

type S1 struct {
	name string
	age  int
}

type S2 struct {
	name string
	age  int
}

type S3 struct {
	hellow string
	S2
}

type hh interface {
	SetName(s string)
	GetName() string
}

type hhs struct {
	name string
	age  int
	sex  int
}

func (h *hhs) SetName(s string) {
	h.name = s
}

func (h *hhs) GetName() string {
	return h.name
}

func (s S2) String() string {
	return fmt.Sprintf("s2 name is %s. age is %d", s.name, s.age)
}

func (s S3) String() string {
	return fmt.Sprintf("s3 said: %s", s.hellow)
}

var sinTest string

type oncx struct {
	m sync.Once
}

var onc oncx

func singleton() {
	onc.m.Do(func() {

	})
}

func main() {

	fmt.Println(sinTest)

	ch := make(chan int)

	go func() {
		<-ch
		fmt.Println("==")
		time.Sleep(time.Second * 5)
		_, _ = http.Get("http://0.0.0.0:8089")
	}()
	time.Sleep(time.Second * 3)

	return

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("i=", i)
			wg.Done()
		}(i)
	}

	err := wg.Wait
	if err != nil {
		fmt.Println("=====", err)
	}

	return

	sqlStr := fmt.Sprintf("select %s from %s where buyer_id = %d and batch_num = '%s'", "111", "www", 112, ";!@#$$##%^^&***()")
	fmt.Printf(sqlStr)
	return

	h := &hhs{}
	h.SetName("hlj")
	fmt.Printf(h.GetName())

	return
	s := S2{name: "zhangsan", age: 15}
	s3 := S3{hellow: "hello world", S2: s}
	fmt.Printf("hahahahh: %s", s3)
	return
	xxxx := [4]int{1, 2, 3, 4}

	fct1 := func(x [4]int) int {
		x[2] = 333
		return x[2]
	}

	fmt.Println(fct1(xxxx), xxxx)

	return

	// s1 := interface{}(&S1{
	// 	"aa", 1,
	// })src/algo/josephus.go
	//
	// switch xx := s1.(type) {
	// case *S1:
	// 	xx.name = "hlj"
	// 	fmt.Println("s1", xx)
	// case *S2:
	// 	xx.name = "s2"
	// 	fmt.Println("s2", xx)
	//
	// default:
	// 	fmt.Println("xxxx", xx)
	// }
	//
	// return
	// name := map[int]int{
	// 	1: 1,
	// 	2:2,
	// }
	// // var name interface{}
	// // name = 22
	//
	// fmt.Println(interface{}(name).(int))
	//
	// v, ok := interface{}(name).(int)

	// v[1] = "ssss"
	// fmt.Println(v, ok, interface{}(name), reflect.TypeOf(v).Elem())

	// name := [5]int{1,2,3,4,5}
	// name1 := name[1:]
	// fmt.Println(len(name1), cap(name1), name1)

	name := make(chan int, 1)
	name <- 1
	close(name)
goto1:
	for {
		select {
		case _, ok := <-name:
			if !ok {
				break goto1
				name = make(chan int)
			}
			fmt.Println("========")
		default:
			fmt.Println("nnnnnnnnn")
			goto goto2
		}
	}

goto2:
	fmt.Println("===goto2")

	return

	t := Teacher{}
	t.ShowA()
	return
	runtime.GOMAXPROCS(1)
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	return
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")

	var i byte
	go func() {
		for i = 0; i <= 255; i++ {
		}
	}()
	fmt.Println("Dropping mic")
	// Yield execution to force executing other goroutines
	runtime.Gosched()
	runtime.GC()
	fmt.Println("Done")

	return

	stu := &Student{
		number:   1,
		realname: "王小明",
		age:      18,
	}
	fmt.Println(stu)

	return

	zhoujielun(&student{
		Name: "zhangsan",
	})

	return
	fmt.Println("normally returned from main")
	//
	// fmt.Errorf("sss")
	//
	// fmt.Println(time.Now().Unix())
	//
	// ch := make(chan int)
	// go producer(ch)
	// for v := range ch {
	// 	fmt.Println("Received ",v)
	// }
	//
	// done := make(chan bool)
	// go hello(done)
	// <-done
	// fmt.Println("main function")

	// a := time.Now()
	//
	// time.Sleep(time.Second * 2)
	// c := time.Now()
	// time.Sleep(time.Second * 2)
	//
	// b :=  time.Now()
	// time.Sleep(time.Second * 2)
	// d :=  time.Now()
	//
	// if b.Before(a) {
	// 	println("b < a")
	// } else {
	// 	println("a < b")
	// }
	//
	// if c.Before(b) {
	// 	println("c < b")
	// }
	//
	// if d.After(b) {
	// 	println("d > b")
	// }

	// go f()
	// c <- 0
	// print(a)

	// fmt.Println(BizGroup)
	//
	// // t := 14 & 1
	// // fmt.Println("ss", t)
	//
	// var ch = make(chan int)
	//
	// go func() {
	// 	fmt.Println("ggg")
	// 	<-ch
	// 	ch <- 2
	// 	ch <- 3
	// 	ch <- 4
	// }()
	//
	// ch <- 1
	// fmt.Println(<- ch)
	// time.Sleep(time.Second * 10)
	// fmt.Println(<- ch)
	// // fmt.Println(<- ch)
	//
	// // go func() {
	// // 	x,_ :=	<-ch
	// // 	fmt.Println("sss", x)
	// // }()
	// // ch <- 1

	// var c = make(chan int)
	// c <- 1
	//
	// close(c)
	//
	// x, ok := <-c
	// // y, ok1 := <-c
	// // z, ok2 := <-c
	//
	//
	// fmt.Println(ch, c, x,ok)
	//
	// // time.Sleep(10*time.Minute)
	//
	// var a float64 = 0
	// var b float64 = 0
	// fmt.Println(a/b)
	// a := []int{1, 3, 4}
	//
	// b := a[1:2:2]
	//
	// fmt.Printf("%d, %d", len(b), cap(b))
	// fmt.Printf("%p, %v, %p, %v", &a[2],   a, b, b)
	//
	// fmt.Println(AA, BB, A, B, E)
	// var s interface{} = map[int]int{1:1}
	// assert(s)
	//
	// a := make(map[int]int)
	//
	// a[1] = 1
	//
	// x := a
	// a[1] = 2
	//
	// fmt.Printf("%p, %p", a, x)
	//
	// b := sync.Map{}
	//
	// b.Store("a", 1)
	// b.Store("b", "sss")
	// b.Range(func(k, v interface{}) bool {
	// 	fmt.Println(k, v)
	// 	return true
	// })
	//
	// b.Delete("a")
	//
	// age, ok := b.Load("a")
	// fmt.Println(age, ok)
	//
	// fmt.Println(a, b)

}

// type entry struct {
// 	p unsafe.Pointer
// }
// type mapss struct {
// 	mu sync.Mutex
// 	read atomic.Value
// 	dirty map[interface{}]*entry
// 	misses int
// }
//
// type readonly struct {
// 	m map[interface{}]*entry
// 	amended bool
// }

// // sync.Map 适用于什么场景？
// // // map并发读写
// // sync.Map 的内部结构是怎样的？
// type entry struct {
// 	p unsafe.Pointer
// }
// type xxmap struct {
// 	mu sync.Mutex
// 	read atomic.Value
// 	dirty map[interface{}]*entry
// 	misses int
// }
// sync.Map.dirty 什么时候升为read？
// misses > len(dirty)
// sync.Map.misses 的作用是什么？
// misses用来判断当前key是不是最新的
// sync.Map.misses 什么时候进行累加？
// 当指定key在read中不存在 && amended = true累加
// sync.Map.read  和 sync.Map.dirty 的底层结构有什么区别？
// read 多了个amended
// sync.Map.read.amended 的作用是什么？
// 用来表示dirty中是否包含read.m中不存在的数据
