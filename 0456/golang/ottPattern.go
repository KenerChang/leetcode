package ottPattern

func find132pattern(nums []int) bool {
	//
	mins := make([]int, len(nums))
	mins[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		mins[i] = min(mins[i-1], nums[i])
	}

	// iterate nums from tail to head
	stack := []int{}
	for j := len(nums) - 1; j >= 0; j-- {
		if nums[j] <= mins[j] {
			continue
		}

		// pop elements in stack which are less than mins[j]
		for len(stack) > 0 && stack[len(stack)-1] <= mins[j] {
			stack = stack[0 : len(stack)-1]
		}

		if len(stack) > 0 && nums[j] > stack[len(stack)-1] {
			return true
		}

		stack = append(stack, nums[j])
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
