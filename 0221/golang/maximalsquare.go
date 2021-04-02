package maximalsquare

import "math"

func maximalSquare(matrix [][]byte) int {
	// one := '1'
	maxSquare := 0

	// allocate dp
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				size := 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
				dp[i][j] = size

				if size > maxSquare {
					maxSquare = size
				}
			}
		}
	}

	return maxSquare * maxSquare
}

func min(nums ...int) int {
	result := math.MaxInt32
	for _, num := range nums {
		if num < result {
			result = num
		}
	}
	return result
}
