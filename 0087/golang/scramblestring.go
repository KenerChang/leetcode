package scramblestring

func isScramble(s1 string, s2 string) bool {
	n := len(s1)
	dp := make([][][]bool, n+1)

	// init dp
	for i := range dp {
		dp[i] = make([][]bool, n)

		for j := range dp[i] {
			dp[i][j] = make([]bool, n)
		}
	}

	// length 0 is not used, start from length 1
	// i: s1 character index, j : s2 character index
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[1][i][j] = s1[i] == s2[j]
		}
	}

	for length := 2; length <= n; length++ {
		// i: s1 start from, j: s2 start from
		for i := 0; i < n+1-length; i++ {
			for j := 0; j < n+1-length; j++ {
				for newLength := 1; newLength < length; newLength++ {
					dp1 := dp[newLength][i]                  // first part
					dp2 := dp[length-newLength][i+newLength] // second part

					// not swap
					dp[length][i][j] = dp[length][i][j] ||
						(dp1[j] && dp2[j+newLength])

					// swap
					dp[length][i][j] = dp[length][i][j] ||
						(dp1[j+length-newLength] && dp2[j])
				}
			}
		}
	}

	return dp[n][0][0]
}
