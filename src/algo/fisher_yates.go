package algo

import (
	"fmt"
	"math/rand"
	"time"
)

/**
	fisher-yates shuffle洗牌算法 有限集合随机等概率打散
	原理：
	第1次 从[0, 7)范围内随机选一个数，拿出，并从候选集合中剔除
	弟2次， 从[0,6)范围内随机选一个数，拿出，并从候选集合中剔除
	。。。。。
	依次类推， 直到所有的数都拿完
 */
func main()  {
	arr := []int{1,2,3,4,5,6,7,8}

	rand.Seed(time.Now().UnixNano())

	for i := len(arr) - 1; i > 0 ; i-- {
		idx := rand.Intn(i)
		arr[i], arr[idx] = arr[idx], arr[i]
	}

	fmt.Println(arr)
}
