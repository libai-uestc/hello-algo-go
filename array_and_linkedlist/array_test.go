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
	fmt.Println("在 nums 中获取随机元素", randomNum)
}
