package jumpgamevi

func maxResult(nums []int, k int) int {
	// we can use dp & deque to solve this problem
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	dq := []int{
		0,
	}

	for i := 1; i < len(nums); i++ {
		for len(dq) > 0 && i-dq[0] > k {
			dq = dq[1:]
		}

		dp[i] = dp[dq[0]] + nums[i]

		for len(dq) > 0 && dp[dq[len(dq)-1]] < dp[i] {
			dq = dq[:len(dq)-1]
		}
		// fmt.Printf("dq: %v\n", dq)
		dq = append(dq, i)
	}

	// fmt.Printf("dp: %v\n", dp)

	return dp[len(dp)-1]
}
