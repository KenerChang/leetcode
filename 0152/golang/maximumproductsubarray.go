package maximumproductsubarray

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	var maxProduct, cmax, cmin int = nums[0], nums[0], nums[0]

	for _, num := range nums[1:] {
		if num < 0 {
			cmax, cmin = cmin, cmax
		}

		cmax = max(cmax*num, num)
		cmin = min(cmin*num, num)
		if cmax > maxProduct {
			maxProduct = cmax
		}
	}
	return maxProduct
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
