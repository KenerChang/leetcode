package besttimetobuyandsellstockiii

import "fmt"

func maxProfit(prices []int) int {
	maxProfitFromStart := make([]int, len(prices))
	minBuyPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		maxProfitFromStart[i] = max(maxProfitFromStart[i-1], prices[i]-minBuyPrice)
		minBuyPrice = min(minBuyPrice, prices[i])
	}

	maxProfitToEnd := make([]int, len(prices)+1)
	maxSellPrice := prices[len(prices)-1]

	for i := len(prices) - 2; i >= 0; i-- {
		maxProfitToEnd[i] = max(maxProfitToEnd[i+1], maxSellPrice-prices[i])
		maxSellPrice = max(maxSellPrice, prices[i])
	}

	fmt.Printf("maxProfitFromStart: %v, maxProfitToEnd: %v\n", maxProfitFromStart, maxProfitToEnd)

	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		maxProfit = max(maxProfit, maxProfitFromStart[i]+maxProfitToEnd[i+1])
	}

	return maxProfit
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
