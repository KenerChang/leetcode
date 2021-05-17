package maximumsubarray

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i], nums[i]+nums[i-1])

		if nums[i] > maxSum {
			maxSum = nums[i]
		}
	}

	return maxSum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
