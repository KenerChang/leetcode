package paintfence

func numWays(n int, k int) int {
	if k == 1 {
		if n > 2 {
			return 0
		} else {
			return 1
		}
	}
	// if n == 1, there are k different colors can pant
	if n == 1 {
		return k
	}

	// if n == 2 , there are k type of same colors and k * k-1 type of 2 colors
	if n == 2 {
		return k * k
	}

	// use dynamic programming
	// there two situations
	// 1. we choose same color as position[i-1] => (k-1) * dp[i-2]
	// 2. we choose different color from position[i-1] => (k-1) * dp[i-1]
	// transform formula is dp[i] = (k-1) * (dp[i-1]+dp[i-2])
	dp := make([]int, n)
	dp[0] = k
	dp[1] = k * k

	for i := 2; i < n; i++ {
		dp[i] = (k - 1) * (dp[i-1] + dp[i-2])
	}

	return dp[n-1]
}
