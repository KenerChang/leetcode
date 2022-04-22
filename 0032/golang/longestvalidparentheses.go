package longestvalidparentheses

func longestValidParentheses(s string) int {
	// init a slice to keep lonest VP of s[i]
	dp := make([]int, len(s))

	// iterate s
	maxVal := 0
	for i, char := range s {
		if i == 0 {
			dp[i] = 0
			continue
		}

		// if s[i] == '(', dp[i] == 0
		if char == '(' {
			dp[i] = 0
		} else {
			// find position of matched '(' j => i -1 -dp[i-1]
			j := i - 1 - dp[i-1]

			// if j < 0, no matched '(', dp[i] = 0
			if j < 0 {
				dp[i] = 0
			} else {
				// else
				// if s[j] == ')', no matched '(', dp[i] = 0
				if s[j] == ')' {
					dp[i] = 0
				} else {
					// get lastVP
					lastVP := 0
					if j >= 1 {
						lastVP = dp[j-1]
					}

					// get innerVP
					innerVP := 0
					if i >= 1 {
						innerVP = dp[i-1]
					}

					// dp[i] = lastVP + innerVP + 2
					dp[i] = lastVP + innerVP + 2

					if dp[i] > maxVal {
						maxVal = dp[i]
					}
				}
			}
		}
	}

	return maxVal
}
