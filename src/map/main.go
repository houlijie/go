package main

import "fmt"

func main() {
	map1 := map[string]string{
		"aaa": "aa",
	}

	fmt.Println(len(map1))

	map2 := map1
	map2["bbbb"] = "bb"
	map1["dd"] = "dd"
	//map[aaa:aa bbbb:bb]   因为map是引用类型，当map被赋值为一个新变量时， 它们指向同一个内部结构
	// 改变其中一个变量 会影响另外一个变量
	fmt.Println(map1) //map[aaa:aa bbbb:bb dd:dd]
	fmt.Println(map2) //map[aaa:aa bbbb:bb dd:dd]

	// map不能使用===判断， == 只能用来检查map是否是nil
	// if map1 == map2 { // err: map can only be compared to nil
	// 	fmt.Println("map1 == map2")
	// }
	if map1 == nil {
		fmt.Println("map1 is nil")
	} else {
		fmt.Println("map1 is not nil")
	}

	map3 := make(map[string]int)
	fmt.Println(map3)

	// map切片
	mapSlice1 := make([]map[string]int, 2)
	// mapSlice1[0]["1"] = 100 // err 没有初始化map, assignment to entry in nil map
	mapSlice1[0] = make(map[string]int, 2)
	mapSlice1[0]["aa"] = 1
	fmt.Println(mapSlice1)
}
