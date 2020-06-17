package maximumproductsubarray

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	var max, c int
	for i, num := range nums {
		c = num
		if c > max {
			max = c
		}

		for j := i + 1; j < len(nums); j++ {
			c = c * nums[j]
			if c > max {
				max = c
			}
		}
	}

	return max
}
