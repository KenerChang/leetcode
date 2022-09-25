package cheapestflightswithinkstops

import (
	"fmt"
	"math"
)

var (
	outdegrees map[int][][]int = map[int][][]int{}
	cache      map[string]int  = map[string]int{}
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// We can solve this problem by using DFS with prune.

	// use a map to keep directed flights and prices
	outdegrees = map[int][][]int{}
	cache = map[string]int{}
	for _, flight := range flights {
		if _, ok := outdegrees[flight[0]]; !ok {
			outdegrees[flight[0]] = [][]int{}
		}
		outdegrees[flight[0]] = append(outdegrees[flight[0]], flight[1:])
	}

	result := dfs(src, dst, k)

	if result == math.MaxInt {
		return -1
	}

	return result
}

func dfs(src, dst, k int) int {
	if src == dst {
		return 0
	}

	if k < 0 {
		return math.MaxInt
	}

	key := fmt.Sprintf("%d-%d", src, k)
	if _, ok := cache[key]; ok {
		return cache[key]
	}

	var minCost int = math.MaxInt
	for _, flight := range outdegrees[src] {
		cost := dfs(flight[0], dst, k-1)

		if cost != math.MaxInt {
			minCost = min(minCost, cost+flight[1])
		}
	}

	cache[key] = minCost
	// fmt.Printf("src: %d, dst: %d, level: %d, minCost: %d\n", src, dst, k, minCost)
	return minCost
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
