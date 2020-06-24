package findpeakelement

import (
	"math"
)

func findPeakElement(nums []int) int {
	min := math.MinInt32
	size := len(nums)

	var prev, next int
	for i, num := range nums {
		if i == 0 {
			prev = min
		} else {
			prev = nums[i-1]
		}

		if i == size-1 {
			next = min
		} else {
			next = nums[i+1]
		}

		if num > prev && num > next {
			return i
		}

	}
	return 0
}
