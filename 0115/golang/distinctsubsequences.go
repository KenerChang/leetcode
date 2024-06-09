package distinctsubsequences

func numDistinct(s string, t string) int {
	// use a matrix to store result
	dp := make([][]int, len(t))

	for i := 0; i < len(t); i++ {
		dp[i] = make([]int, len(s))
	}

	for i := 0; i < len(t); i++ {
		for j := i; j < len(s); j++ {
			if j == i {
				if s[j] == t[i] {
					if i == 0 {
						dp[i][j] = 1
					} else {
						dp[i][j] = dp[i-1][j-1]
					}
				}
			} else {
				dp[i][j] = dp[i][j-1] // not use s[j]
				if s[j] == t[i] {
					// use s[j]
					if i == 0 {
						dp[i][j] += 1
					} else {
						dp[i][j] += dp[i-1][j-1]
					}
				}
			}
		}
	}

	return dp[len(t)-1][len(s)-1]
}
