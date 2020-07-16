package houserobber

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	for i, num := range nums {
		if i == 0 {
			continue
		} else if i == 1 {
			nums[i] = max(nums[0], nums[1])
		} else {
			nums[i] = max(nums[i-2]+num, nums[i-1])
		}
	}
	return nums[len(nums)-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
