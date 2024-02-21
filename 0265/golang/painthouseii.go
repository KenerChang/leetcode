package painthouseii

import (
	"math"
)

func minCostII(costs [][]int) int {
	// we can iterate costs matrix
	// for each house, we keep the mininum 2 cost and of the house with its position
	// And for each costs cell costs[i][j], min_costs[i][j] = costs[i][j] + min available cost from last house

	var mininum_costs [][]int
	var last_house_mininum_costs [][]int
	for i := 0; i < len(costs); i++ {
		last_house_mininum_costs = mininum_costs
		mininum_costs = [][]int{{math.MaxInt, -1}, {math.MaxInt, -1}}
		for j := 0; j < len(costs[i]); j++ {
			if i > 0 {
				if j == last_house_mininum_costs[0][1] {
					costs[i][j] += last_house_mininum_costs[1][0]
				} else {
					costs[i][j] += last_house_mininum_costs[0][0]
				}
			}

			if costs[i][j] <= mininum_costs[0][0] {
				mininum_costs[1] = mininum_costs[0]
				mininum_costs[0] = []int{costs[i][j], j}
			} else if costs[i][j] < mininum_costs[1][0] {
				mininum_costs[1] = []int{costs[i][j], j}
			}
		}

		// fmt.Printf("costs: %v, mininum_costs: %v\n", costs, mininum_costs)
	}

	return mininum_costs[0][0]
}
