package subarraysumequalsk

func getSum(i, j int, nums []int) int {
	// get sum of nums[i:j]
	if i == j {
		if i == 0 {
			return nums[0]
		} else {
			return nums[i] - nums[i-1]
		}
	} else {
		if i == 0 {
			return nums[j]
		} else {
			return nums[j] - nums[i-1]
		}
	}
}

func subarraySum(nums []int, k int) int {
	// try sum up nums
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	var count int
	for i := 0; i < len(nums); i++ {
		for j := 0; i+j < len(nums); j++ {
			if getSum(i, i+j, nums) == k {
				count++
			}
		}
	}
	return count
}
