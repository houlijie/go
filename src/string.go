package main

import (
	"fmt"
	"math"
)

// 根据offset和间隔获取每页插入数据的位置
// func getLastPositionRangeByOffset(offset int64, pageSize int64, interval int64) map[int64]int64 {
// 	postionMap := make(map[int64]int64)

// 	lastOffset := offset + pageSize                                              // items总数                                                                         // 总数
// 	total := math.Floor(float64(lastOffset) / float64(interval))                 // 穿插总数量
// 	currentPageNum := int(total - math.Floor(float64(offset)/float64(interval))) // 当前页穿插数量 = 穿插总数量 - 历史穿插总个数

// 	lastPosition := int(total) * int(interval)               // 最后一个穿插的位置
// 	lastIndex := int(math.Mod(total, float64(interval))) - 1 // 卡片类型index

// 	for i := 0; i < currentPageNum; i++ {
// 		posID := int64(lastPosition - i*int(interval))
// 		CardTypeIndex := int64(lastIndex - i)
// 		postionMap[posID] = CardTypeIndex
// 	}

// 	return postionMap
// }

// func GetRangeItemsByOffset(start int64, end int64, step int64, items []string, firstItemPos int64) map[int64]string {
// 	rangeMap := make(map[int64]string)
// 	if start < 0 || end <= 0 || step <= 0 {
// 		return rangeMap
// 	}

// 	if end < firstItemPos {
// 		return rangeMap
// 	}

// 	var lastItemPos, lastItemCount int64
// 	if start >= firstItemPos {
// 		fmt.Println(start - firstItemPos)
// 		lastItemCount = int64(math.Floor(float64(start-firstItemPos+step) / float64(step))) // 已出现item个数
// 		lastItemPos = lastItemCount*step + firstItemPos - step                              // 上一次出现的位置
// 	} else {
// 		lastItemPos = firstItemPos - step // 0的位置
// 	}

// 	fmt.Println("lastItemCount:", lastItemCount, "lastItemPos", lastItemPos)

// 	itemsTotal := int64(math.Floor(float64(end-firstItemPos+step) / float64(step))) // 插入总次数
// 	fmt.Println("itemsTotal", itemsTotal)
// 	rangeItemCount := itemsTotal - lastItemCount
// 	var i int64
// 	itemsCount := len(items)
// 	for ; i < rangeItemCount; i++ {
// 		itemPos := lastItemPos + step*(i+1)
// 		itemIndex := int(math.Mod(float64(lastItemCount+i), float64(itemsCount)))
// 		rangeMap[itemPos] = items[itemIndex]
// 	}

// 	return rangeMap
// }
func GetRangeItemsByOffset(start int64, end int64, step int64, items []string, firstItemPos int64) map[int64]string {
	rangeMap := make(map[int64]string)
	if start < 0 || end <= 0 || step <= 0 {
		return rangeMap
	}

	if end < firstItemPos {
		return rangeMap
	}

	var lastItemPos, lastItemCount int64
	if start >= firstItemPos {
		lastItemCount = int64(math.Floor(float64(start-firstItemPos+step) / float64(step))) // 已出现item个数
		lastItemPos = lastItemCount*step + firstItemPos - step                              // 上一次出现的位置
	} else {
		lastItemPos = firstItemPos - step // 0的位置
	}

	itemsTotal := int64(math.Floor(float64(end-firstItemPos+step) / float64(step))) // 插入总次数
	rangeItemCount := itemsTotal - lastItemCount
	var i int64
	itemsCount := len(items)
	for ; i < rangeItemCount; i++ {
		itemPos := lastItemPos + step*(i+1)
		itemIndex := int(math.Mod(float64(lastItemCount+i), float64(itemsCount)))
		rangeMap[itemPos] = items[itemIndex]
	}

	return rangeMap
}

type ChannelConf struct {
	ChannelType string `json:"channel_type"`
	Id          int    `json:"id,string"`
}

type extra struct {
	Id string
}

func main() {
	ahhaha := make(map[int]int, 1)
	test := []int{
		1,
	}

	fmt.Println(ahhaha)
	for _, mallId := range test {
		ahhaha[mallId] = mallId
	}
	fmt.Println(ahhaha)

	return
	// cardsMap[5] = 111115
	// cardsMap[10] = 10000
	// cardsMap[15] = 15555
	// cardsMap[20] = 2000000
	// cardsMap[25] = 200220000
	// cardsMap[45] = 111115
	// cardsMap[50] = 10000
	// cardsMap[55] = 15555
	// cardsMap[60] = 2000000
	// cardsMap[65] = 200220000

	// goods := make([]int, 0, 0)

	// for i := 1; i <= 20; i++ {
	// 	goods = append(goods, i)
	// }

	// fmt.Println(goods)

	// posList := make([]int, 0, 0)
	// for pos := range cardsMap {
	// 	fmt.Println(pos)
	// 	posList = append(posList, int(pos))
	// }

	// fmt.Println(posList)
	// sort.Ints(posList)
	// fmt.Println(posList)

	// for _, pos := range posList {
	// 	card, _ := cardsMap[int64(pos)]
	// 	pos = pos - 40
	// 	if pos > len(goods) {
	// 		fmt.Println("goods", pos, len(goods))
	// 		goods = append(goods, card)
	// 	} else {
	// 		index := int(pos) - 1
	// 		fmt.Println("index", index)
	// 		goods = append(goods[:index], append([]int{card}, goods[index:]...)...)
	// 	}
	// }

	// fmt.Println(goods)
	// return

	// return

	// test1, _ := json.Marshal(extra{
	// 	Id: "3519",
	// })

	// var testxxxx ChannelConf
	// err := json.Unmarshal(test1, &testxxxx)

	// fmt.Println(testxxxx, err)
	// return
	// items := make(map[int]int)
	// offset := 20
	// for i := 0; i < 20; i++ {
	// 	// 穿插卡片
	// 	position := (i + 1) + offset
	// 	items[position] = position
	// }

	// fmt.Println(items)
	// return

	// aaa := "a"
	// bbb := "b"

	// ddd := aaa + bbb
	// fmt.Println("ssss", ddd)
	// cccc := &ddd{}
	// fmt.Printf("ssss %v", cccc)
	// return

	a := ""
	// fmt.Printf("%+v", len(a))

	// return

	abc := []string{
		"a", "b", "c", "d",
	}
	test := GetRangeItemsByOffset(0, 20, 5, abc, 30)
	// for posId, cardType := range test {
	// 	fmt.Println(posId, cardType)
	// }
	fmt.Println("%+v", test)
	return
	aa := []string{
		"aa",
	}
	var bb int = 0
	for i := 0; i < 3; i++ {
		dd := bb - i
		fmt.Printf("aaa: \n", aa[dd])
	}
	return
	x := float64(5)
	y := float64(4)
	t := math.Mod(x, y)
	fmt.Printf("xxxxx%T%v   ", t, t)
	return
	aaaa := map[string]string{
		"gId": "gId",
	}
	_, ok := aaaa["gid"]
	fmt.Println("aaa%v", ok)
	return

	//
	// a = "hello world"

	// go字符串类型不能改变
	// a = false //  cannot use false (type bool) as type string in assignment
	// a = "bbb" // bbb
	fmt.Println(a)

	// 双引号可以识别转义字符 反引号``原样输出， 可以实现防攻击

	// var b = "aa \n bb \n"
	// var c = `aa \n bb \n`
	// fmt.Println(b, c)

	// 字符串的拼接 +, 当一个拼接的操作很长时，可以分行写， +号必须在上一行
	str1 := "hello"
	str2 := "world"
	d := str1 + str2
	fmt.Println(d) // helloworld
	d += " everyone"
	fmt.Println(d) // helloworld everyone

	d = "hello" + "hello" +
		"hello" + "hello" +
		"hello"
	fmt.Println(d)
}
