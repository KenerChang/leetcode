package subsetsii

import "fmt"
import "sort"

func subsetsWithDup(nums []int) [][]int {
	// dynamic programming for all subsets
	// but should filter duplicate slices
	if len(nums) == 0 {
		return [][]int{}
	}
	results := [][]int{
		[]int{},
	}

	cache := map[string]bool{}

	for _, num := range nums {
		for _, result := range results {
			newResult := append([]int{}, result...)
			newResult = append(newResult, num)

			sort.Ints(newResult)
			resultKey := fmt.Sprint(newResult)

			if _, ok := cache[resultKey]; !ok {
				results = append(results, newResult)
				cache[resultKey] = true
			}
		}
	}
	return results
}
