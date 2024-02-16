package burstballoons

func maxCoins(nums []int) int {
	// Instead of considersing which balloon to burst first,
	// we consider which balloon to burst last.

	// add 1 to head and tail of nums
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	// init db table
	table := make([][]int, len(nums))
	for i := 0; i < len(table); i++ {
		table[i] = make([]int, len(nums))
	}

	return dp(nums, 1, len(nums)-2, table)
}

func dp(nums []int, l, r int, table [][]int) int {
	if r < l {
		return 0
	}

	if table[l][r] > 0 {
		return table[l][r]
	}

	var result int
	for i := l; i <= r; i++ {
		gain := nums[l-1] * nums[i] * nums[r+1]
		remaining := dp(nums, l, i-1, table) + dp(nums, i+1, r, table)

		result = max(gain+remaining, result)
	}

	table[l][r] = result
	return result
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
