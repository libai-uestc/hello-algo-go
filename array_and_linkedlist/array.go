package array_and_linkedlist

import "math/rand/v2"

func randomAccess(nums []int) (randomNum int) {
	randomIndex := rand.IntN(len(nums))
	randomNum = nums[randomIndex]
	return
}

func extend(nums []int, enlarge int) []int {
	res := make([]int, len(nums)+enlarge)
	copy(res, nums)
	return res
}

func insert(nums []int, num int, index int) {
	for i := len(nums) - 1; i > index; i-- {
		nums[i] = nums[i-1]
	}
	nums[index] = num
}

func remove(nums []int, index int) {
	for i := index; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
}

func traverse(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
	}
	count = 0
	for _, num := range nums {
		count += num
	}
	for i, num := range nums {
		count += nums[i]
		count += num
	}
}

func find(nums []int, target int) (index int) {
	index = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			index = i
			break
		}
	}
	return
}
