package triangle

// import (
// 	"math"
// )

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	if len(triangle) == 1 {
		return triangle[0][0]
	}

	for rowIdx := len(triangle) - 2; rowIdx >= 0; rowIdx-- {
		for columnIdx := 0; columnIdx <= rowIdx; columnIdx++ {
			min := triangle[rowIdx+1][columnIdx]
			if triangle[rowIdx+1][columnIdx+1] < min {
				min = triangle[rowIdx+1][columnIdx+1]
			}

			triangle[rowIdx][columnIdx] = triangle[rowIdx][columnIdx] + min
		}
	}

	return triangle[0][0]
}
