package combination_sum_2
import (
	"sort"
	"fmt"
)

func combinationSum2(candidates []int, target int) [][]int {
	if target <= 0 {
		return [][]int{}
	}

	combinations := [][]int{}

	exist := map[string]bool{}
	for idx, candidate := range candidates {
		if target == candidate {
			combination := []int {candidate}
			key := fmt.Sprint(combination)

			if _, ok := exist[key]; !ok {
				combinations = append(combinations, combination)
				exist[key] = true
			}
			continue
		}

		remain := target - candidate
		if remain > 0 && idx < len(candidates) - 1 {
			candidateCopy := make([]int, len(candidates))
			copy(candidateCopy, candidates)
			leftCandidates := append(candidateCopy[0:idx], candidateCopy[idx+1:]...)
			combinationsOfRemain := combinationSum2(leftCandidates, remain)

			// copy combinations
			for _, combination := range combinationsOfRemain {
				combinationWithCandidate := append(combination, candidate)

				sort.Ints(combinationWithCandidate)
				key := fmt.Sprint(combinationWithCandidate)
				if _, ok := exist[key]; !ok {
					exist[key] = true
					combinations = append(combinations, combinationWithCandidate)
				}
			}
		}
	}
	return combinations
}