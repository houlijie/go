package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	蓄水池抽样算法 集合nums长度为n，无限大，随机等概率获取k个数据

	原理
	1、n < k 直接返回nums
	2、n > k 初始化蓄水池, 在[k, n]范围内取随机值idx，若idx落入[0, k-1]的范围内， 则用nums[i]替换蓄水池中的nums[idx]


*/
func sampling(nums []int, k int) []int  {
	len := len(nums)
	if len < k {
		k = len
	}

	res := make([]int, k, k) // ! len
	copy(res, nums)

	rand.Seed(time.Now().UnixNano())
	for i := k; i < len; i++ {
		idx := rand.Intn(i)
		if idx < k {
			res[idx] = nums[i]
		}
	}

	return res
}

func main()  {
	fmt.Println(sampling([]int{1,3,24,45,52,21,3,4,2}, 3))
}