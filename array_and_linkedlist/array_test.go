package array_and_linkedlist

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var arr [5]int
	fmt.Println("数组 arr =", arr)
	nums := []int{1, 3, 2, 5, 4}
	fmt.Println("数组 nums =", nums)

	randomNum := randomAccess(nums)
	fmt.Println("在 nums 中获取随机元素，获取到的元素是：", randomNum)

	nums = extend(nums, 3)
	fmt.Println("将数组长度扩展至 8 , 得到 nums =", nums)

	insert(nums, 6, 3)
	fmt.Println("在索引 3 处插入数字 6 , 得到 nums =", nums)

	remove(nums, 2)
	fmt.Println("删除索引 2 处的元素, 得到 nums =", nums)

	traverse(nums)

	index := find(nums, 3)
	fmt.Println("在 nums 中查找元素 3 , 得到索引 =", index)
}

// go test -v .\hello-algo-go\array_and_linkedlist\ -run=^TestArray$ -count=1
