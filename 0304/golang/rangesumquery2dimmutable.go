package rangesumquery2dimmutable

type NumMatrix struct {
	dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	// dp[i][j] means sum of (0, 0) to (i, j)

	dp := make([][]int, len(matrix)+1)
	for i := 0; i <= len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0])+1)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			// dp[r + 1][c] + dp[r][c + 1] + matrix[r][c] - dp[r][c];
			dp[i+1][j+1] = dp[i+1][j] + dp[i][j+1] + matrix[i][j] - dp[i][j]
		}
	}

	return NumMatrix{
		dp: dp,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.dp[row2+1][col2+1] - this.dp[row2+1][col1] - this.dp[row1][col2+1] + this.dp[row1][col1]
}
