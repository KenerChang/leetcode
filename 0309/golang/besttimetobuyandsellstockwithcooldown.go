package besttimetobuyandsellstockwithcooldown

import (
	"math"
)

func maxProfit(prices []int) int {
	// we have 3 states here: sold, held, and reset
	// and 3 actions: sale, hold, rest.

	// For sold, there is only 1 action to perform: rest.
	// For held, there are 2 actions to perform: rest, sale
	// For reset, there are 2 actions to perform: rest, hold

	// So for each price index i, we ca get max profit by max(sold[i], held[i], reset[i])
	// and to get sold profit, the formula is held[i-1] + price
	// ,to get held profit, the formula is max(held[i-1], reset[i]-1 - price)
	// and to get rest profit, the formula is max(reset[i-1], sold[i-1])

	soldMaxProfit := math.MinInt
	heldMaxProfit := math.MinInt
	resetMaxProfit := 0

	for i := 0; i < len(prices); i++ {
		soldMaxProfitForI := heldMaxProfit + prices[i]

		heldMaxProfitForI := max(heldMaxProfit, resetMaxProfit-prices[i])

		resetMaxProfitForI := max(resetMaxProfit, soldMaxProfit)

		soldMaxProfit = soldMaxProfitForI
		heldMaxProfit = heldMaxProfitForI
		resetMaxProfit = resetMaxProfitForI
	}

	return max(soldMaxProfit, heldMaxProfit, resetMaxProfit)
}

func max(nums ...int) int {
	var result int = nums[0]

	for i := 1; i < len(nums); i++ {
		if result < nums[i] {
			result = nums[i]
		}
	}

	return result
}
