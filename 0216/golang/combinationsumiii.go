package combinationsumiii

import (
// "fmt"
)

func combinationSum3(k int, n int) [][]int {
	results := [][]int{}
	if k == 0 || k >= 10 {
		return results
	}

	total := 0
	for i := 10 - k; i < 10; i++ {
		total += i
	}

	if total < n {
		return results
	}

	digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	combinationSum3Recurive(k, n, 0, digits, []int{}, &results)
	return results
}

func combinationSum3Recurive(k, n, start int, digits []int, result []int, results *[][]int) {
	if k <= 0 || n == 0 {
		return
	}

	// fmt.Printf("k: %d, n: %d, result: %v\n", k, n, result)

	for i := start; i < len(digits); i++ {
		digit := digits[i]
		if k == 1 && n-digit == 0 {
			// a candidate
			result = append(result, digit)
			*results = append(*results, result)
			return
		} else {
			r := append([]int{}, result...)
			r = append(r, digit)
			combinationSum3Recurive(k-1, n-digit, i+1, digits, r, results)
		}
	}
}
