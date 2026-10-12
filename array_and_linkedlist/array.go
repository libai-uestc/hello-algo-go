package array_and_linkedlist

import "math/rand/v2"

func randomAccess(nums []int) (randomNum int) {
	randomIndex := rand.IntN(len(nums))
	randomNum = nums[randomIndex]
	return
}
