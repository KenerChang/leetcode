package coinchange

func coinChange(coins []int, amount int) int {
	// we can use BFS to find the shortest path of a solution

	if amount == 0 {
		return 0
	}

	queue := [][]int{
		{amount, 0},
	}

	visited := make([]bool, amount)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, coin := range coins {
			if node[0] == coin {
				// a solution return
				return node[1] + 1
			} else if node[0] > coin {
				remain := node[0] - coin

				if visited[remain] {
					continue
				}

				visited[remain] = true
				queue = append(queue, []int{remain, node[1] + 1})
			}
		}
	}

	return -1
}
