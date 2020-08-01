package houserobberii

import (
// "fmt"
)

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) <= 3 {
		maxRob := 0
		for _, num := range nums {
			maxRob = max(maxRob, num)
		}
		return maxRob
	}

	maxRobs := make([]int, len(nums))

	// assume not rob final house
	maxRobs[0] = nums[0]
	maxRobs[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums)-1; i++ {
		maxRobs[i] = max(maxRobs[i-1], maxRobs[i-2]+nums[i])
	}

	c1 := maxRobs[len(nums)-2]

	// assume not rob first house
	maxRobs = make([]int, len(nums))
	maxRobs[1] = nums[1]
	maxRobs[2] = max(nums[1], nums[2])
	for i := 3; i < len(nums); i++ {
		maxRobs[i] = max(maxRobs[i-1], maxRobs[i-2]+nums[i])
	}

	return max(c1, maxRobs[len(nums)-1])
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
