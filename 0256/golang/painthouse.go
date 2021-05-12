package painthouse

import "math"

func minCost(costs [][]int) int {
	if len(costs) == 1 {
		return min(costs[0]...)
	}

	for i := len(costs) - 2; i >= 0; i-- {
		for j := 0; j < len(costs[i]); j++ {
			switch j {
			case 0:
				costs[i][j] += min(costs[i+1][1], costs[i+1][2])
			case 1:
				costs[i][j] += min(costs[i+1][0], costs[i+1][2])
			case 2:
				costs[i][j] += min(costs[i+1][0], costs[i+1][1])
			}
		}
	}

	return min(costs[0]...)
}

func min(nums ...int) int {
	m := math.MaxInt32

	for _, num := range nums {
		if num < m {
			m = num
		}
	}
	return m
}
