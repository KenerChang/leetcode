package besttimetobuyandsellstockii

// import (
// 	"fmt"
// )

func maxProfit(prices []int) int {
	// direction: 1 up, 0 equal, -1 down
	var profit, buyPrice, currDirection int
	for i, price := range prices {
		if i == 0 {
			continue
		}

		direction := price - prices[i-1]
		if direction > 0 && currDirection <= 0 {
			// from down to up
			buyPrice = prices[i-1]
		} else if direction < 0 && currDirection > 0 {
			// from up to down
			profit += prices[i-1] - buyPrice
		}

		if direction != 0 {
			currDirection = direction
		}
	}

	if currDirection > 0 {
		profit += prices[len(prices)-1] - buyPrice
	}
	return profit
}
