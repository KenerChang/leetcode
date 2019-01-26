package combination_sum

import (
	"fmt"
	"sort"
	"encoding/json"
)

func combinationSum(candidates []int, target int) [][]int {
	// use DP and a local cache map to decide if a candidate is duplocated or not
	// time complexity of the implementation is expect to O(mn)
	// m is length of the candidates
	// n is number of target
	if target == 0 {
		return [][]int{}
	}
	combinationsOfTargets := map[int][][]int{}
	candidateMap := map[int]bool{}
	combinationMap := map[string]bool{}

	// create a candidate map which can quickly decide a number is a candidate or not
	for _, candidate := range candidates {
		candidateMap[candidate] = true
	}

	for i :=0; i <= target; i++ {
		combinationsOfi := [][]int{}
		for _, candidate := range candidates {
			remain := i - candidate
			if remain < 0 {
				continue
			}

			if i == candidate {
				combinationsOfi = append(combinationsOfi, []int{candidate})
			}

			if combinationsOfRemain, ok := combinationsOfTargets[remain]; ok {
				combinations := make([][]int, len(combinationsOfRemain))
				// for deep copy
				copyStr, _ := json.Marshal(combinationsOfRemain[:])
				json.Unmarshal(copyStr, &combinations)

				for idx := range combinations {
					combination := append(combinations[idx], candidate)
					sort.Ints(combination)
					key := fmt.Sprint(combination)
					if _, ok := combinationMap[key]; !ok {
						combinationsOfi = append(combinationsOfi, combination)
						combinationMap[key] = true
					}
				}
			}
		}
		combinationsOfTargets[i] = combinationsOfi
	}

	combinations := combinationsOfTargets[target]
	return combinations
}