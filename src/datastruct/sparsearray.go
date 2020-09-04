package main

import "fmt"

type node struct {
	row int
	col int
	val int
}

func eachMap(chessMap [11][11]int) {
	for _, v1 := range chessMap {
		for _, v2 := range v1 {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}

// 稀疏数组sparsearray
/*
	1、记录数组有几行几列，有多少不同的值
	思想： 把具有
*/
func main() {
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	//打印元数据
	eachMap(chessMap)

	//  转成稀疏数组
	var parseArray []node

	// 标准的稀疏数组应该还有一个记录元素的二维数组的规模
	defaultNode := node{
		row: 11,
		col: 11,
		val: 0,
	}
	parseArray = append(parseArray, defaultNode)
	for i, v3 := range chessMap {
		for j, v4 := range v3 {
			if v4 != 0 {
				vNode := node{
					row: i,
					col: j,
					val: v4,
				}
				parseArray = append(parseArray, vNode)
			}
		}
	}

	fmt.Println(parseArray)

	// 恢复元数组
	var chessMap2 [11][11]int
	for i, v := range parseArray {
		if i == 0 {
			continue
		}
		chessMap2[v.row][v.col] = v.val
	}

	eachMap(chessMap2)
}
