package rotatearray

// import (
// 	"fmt"
// )

func rotate(nums []int, k int) {
	k = k % len(nums)

	if k == 0 {
		return
	}

	nLen := len(nums)
	last := nLen - 1

	reverse(nums, 0, last)
	reverse(nums, k, last)
	reverse(nums, 0, k-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
