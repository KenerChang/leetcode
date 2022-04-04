package factorcombinations

import (
	"math"
)

func getFactors(n int) [][]int {
	// for 1 to squre root(n) as k
	// if n % k == 0
	// append result of getFactors(n/k)

	return getFactorsRecursive(n, 2, map[int][][]int{})
}

func getFactorsRecursive(n, start int, cache map[int][][]int) [][]int {
	// if n is handled before, return result in cache
	// if result, ok := cache[n]; ok {
	// 	return result
	// }

	// for 1 to squre root(n) as k
	results := [][]int{}
	for i := start; i <= int(math.Floor(math.Sqrt(float64(n)))); i++ {
		// if n % k == 0
		if n%i != 0 {
			continue
		}

		// append result of getFactors(n/k)
		// results = [][]int{{i, n / i}}
		results = append(results, []int{i, n / i})

		factorsOfI := getFactorsRecursive(n/i, i, cache)
		for _, result := range factorsOfI {
			result = append([]int{i}, result...)
			results = append(results, result)
		}
	}

	cache[n] = results

	return results
}
