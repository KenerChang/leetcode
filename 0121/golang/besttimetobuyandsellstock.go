package besttimetobuyandsellstock

import (
	"math"
)

func maxProfit(prices []int) int {
	max := 0
	curMin := math.MaxInt64
	for _, price := range prices {
		if price < curMin {
			curMin = price
		} else if price-curMin > max {
			max = price - curMin
		}
	}

	return max
}
