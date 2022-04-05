package uglynumberii

func nthUglyNumber(n int) int {
	if n <= 6 {
		return n
	}

	var ptr2, ptr3, ptr5 int
	var results []int = make([]int, n)
	results[0] = 1
	for i := 1; i < n; i++ {
		val2, val3, val5 := results[ptr2]*2, results[ptr3]*3, results[ptr5]*5

		results[i] = min(val2, val3, val5)

		// update ptr
		if results[i] == val2 {
			ptr2++
		}

		if results[i] == val3 {
			ptr3++
		}

		if results[i] == val5 {
			ptr5++
		}
	}

	return results[n-1]
}

func min(nums ...int) int {
	minVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < minVal {
			minVal = nums[i]
		}
	}

	return minVal
}
