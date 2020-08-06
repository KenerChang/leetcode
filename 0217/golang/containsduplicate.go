package containsduplicate

import (
	"sort"
)

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	sort.Ints(nums)
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if prev == num {
			return true
		}
		prev = num
	}

	return false
}
