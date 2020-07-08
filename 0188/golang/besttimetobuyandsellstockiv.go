package besttimetobuyandsellstockiv

import (
	// "fmt"
	"math"
)

func maxProfit(k int, prices []int) int {
	if len(prices) == 0 || k == 0 {
		return 0
	}

	if k > len(prices)/2 {
		// we can be greedy since we have enough transcation limits
		profit := 0
		for i, price := range prices[1:] {
			profit += max(0, price-prices[i])
		}
		return profit
	}

	// use dp, row: transcation count, column: day
	profits := make([][]int, k+1)
	for i := 0; i < len(profits); i++ {
		profits[i] = make([]int, len(prices))
	}

	for t := 1; t < k+1; t++ {
		maxUntilD := math.MinInt32
		for d := 1; d < len(prices); d++ {
			// max profit of transcation count t until d
			maxUntilD = max(maxUntilD, profits[t-1][d-1]-prices[d-1])

			//max profit of transcation count t & day d is one of:
			// 1. do not sell, max profit is same as profits[t][d-1]
			// 2. do sell, max profit is maxUntilD + prices[d]
			profits[t][d] = max(profits[t][d-1], maxUntilD+prices[d])
		}
	}
	return profits[k][len(prices)-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
