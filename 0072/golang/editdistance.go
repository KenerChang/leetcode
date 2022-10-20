package editdistance

func min(nums ...int) int {
	minVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < minVal {
			minVal = nums[i]
		}
	}
	return minVal
}

func minDistance(word1 string, word2 string) int {
	// We can solve this problem with DP.
	// transform fomula:
	// dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1) if word1[i] != word2[j]
	// dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]) else

	// init dp
	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)

		for j := 0; j <= len(word2); j++ {
			if i == 0 {
				dp[i][j] = j // insert j chars to become word2
			} else if j == 0 {
				dp[i][j] = i // delete i chars to become word2
			} else {
				if word1[i-1] == word2[j-1] {
					dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1])
				} else {
					dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
				}
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
