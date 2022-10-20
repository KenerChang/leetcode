package kinversepairsarray

const (
	M = 1000000007
)

func kInversePairs(n int, k int) int {
	// To solve this problem, we can start with k = 0. For example, if n = 3 and k = 0,
	// we have only 1 solution that is [1,2,3] and define it as a0.
	// Start from a0, if k = 1, we have [2,1,3] and [1,3,2], 2 and 3 is move left for 1 position
	// And if we add 4 to the above arrays, we can get [2,1,3,4] and [1,3,2,4] and k is still 2.
	// Now, if we try to move 4 for 2 position left, the arrays become [2,4,3,1] and [1,4,3,2] and k became 4
	// From the above example, we can know count(n, k) = count(n-1, k) + count(n-1, k-2) + ..... + count(n-1, 0)
	// So we can use dp to solve this problem.

	// we can also improve the way we get count(n, k) by applying accumulative sum
	if k == 0 {
		return 1
	}

	dp := make([][]int, n+1)
	dp[0] = make([]int, k+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, k+1)

		// since there are most (i * i - 1) / 2 pairs (reverse sorted)
		// j is min(k, (i * i - 1)/2)
		for j := 0; j <= k; j++ {
			if j == 0 {
				dp[i][j] = 1 // only 1 solution
			} else {
				var val int
				// count(i, j)
				if j >= i {
					// since j >= i, we need count[i-1][0] + ... + count[i-1][i]
					val = (dp[i-1][j] - dp[i-1][j-i]) % M
				} else {
					// k < n, we need count[i-1][0] + ... + count[i-1][k]
					val = (dp[i-1][j] + M) % M
				}
				dp[i][j] = (dp[i][j-1] + val) % M
			}
		}
	}

	return (dp[n][k] + M - dp[n][k-1]) % M
}
