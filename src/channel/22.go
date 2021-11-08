package main

import (
	"fmt"
	"sync"
	"time"
)

var wg10 sync.WaitGroup
// var lock sync.Mutex
//
// // wait 实现wait方法，收集jobs中channel中的int值返回。
// // 每个channel都一定有一个int返回，不会close。
// // 返回slice长度应当与jobs相同，顺序任意。
// func wait(jobs []<-chan int) []int {
// 	res := make([]int, 0, len(jobs))
// 	for _, j := range jobs {
// 		if v, ok := <- j; ok {
// 			// i := v
// 			wg10.Add(1)
// 			go func() {
// 				defer wg10.Done()
// 				lock.Lock()
// 				res = append(res, v)
// 				lock.Unlock()
// 			}()
// 		}
// 	}
//
// 	wg10.Wait()
//
// 	return res
// }

// 以下代码仅为示例、调试用，不用修改
func main() {
	// chd := make(chan int, 3)
	// chd <- 2
	// chd <-3
	// chd <-4
	// // forever := make(chan bool)
	// // // forever <- true
	// // go func() {
	// // 	for {
	// // 		fmt.Println("xxx")
	// // 	}
	// // }()
	// // <- forever
	// return
	xxx := []int{1, 2, 3, 4, 5}

	xxs := xxx[0:0:2]

	// xxs = append(xxs, 10)


	fmt.Println(xxx, xxs, len(xxs), cap(xxs))
	return

	t1 := time.Now().Nanosecond()
	for i := range xxx {
		fmt.Println("range i", xxx[i])
	}
	t2 := time.Now().Nanosecond()
	for _, v := range xxx {
		fmt.Println("range v", v)
	}
	t3 := time.Now().Nanosecond()
	for i := 0; i < len(xxx); i++ {
		fmt.Println("range i", xxx[i])
	}
	t4 := time.Now().Nanosecond()
	for i, v := range xxx {
		fmt.Println("range v",i, v)
	}
	t5 := time.Now().Nanosecond()

	fmt.Println(t2-t1, t3 -t2, t4 - t3, t5 - t4)
	return
	dataMap := sync.Map{}

	fmt.Println(xxx[:3])
	for i, i2 := range xxx {
		wg10.Add(1)
		// i, i2 := i, i2
		i2 := i2
		go func(i int) {
			// fmt.Println(i, i2)
			dataMap.Store(i, i2)
			defer wg10.Done()
		}(i)
	}

	wg10.Wait()

	for x := range xxx {
		data, ok := dataMap.Load(x)

		fmt.Println(data, ok)
	}



	time.Sleep(time.Second* 3)

	return
	// *i=10
	// rand.Seed(1)
	// jobs := begin()
	// results := wait(jobs)
	// /*预期输出例子如下
	// produce 3 518
	// produce 0 625
	// produce 2 656
	// produce 1 740
	// wait[625 740 656 518]
	// */
	// fmt.Print("wait", results)
}

// func begin() (rst []<-chan int) {
// 	for i := 0; i < rand.Intn(10)+3; i++ {
// 		fmt.Println("i=", i)
// 		ch := make(chan int)
// 		ii := i
// 		go func() {
// 			// 模拟一些异步任务
// 			x := rand.Intn(1000) + 200
// 			time.Sleep(time.Millisecond * time.Duration(x))
// 			println("produce", ii, x)
// 			ch <- x
// 		}()
// 		rst = append(rst, ch)
// 	}
//
// 	return
// }
