package combination_sum_2

import (
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	// test candidates = [10,1,2,7,6,1,5], target = 8
	candidates := []int{10,1,2,7,6,1,5}
	target := 8

	result := combinationSum2(candidates, target)
	if len(result) != 4 {
		t.Errorf("test candidates: %+v, target: %d failed", candidates, target)
		t.Errorf("got result: %+v, ", result)
	}

	// test candidates = [2,5,2,1,2], target = 5
	candidates = []int{2,5,2,1,2}
	target = 5

	result = combinationSum2(candidates, target)
	if len(result) != 2 {
		t.Errorf("test candidates: %+v, target: %d failed", candidates, target)
		t.Errorf("got result: %+v, ", result)
	}

	// test candidates = [1,1], target = 1
	candidates = []int{1,1}
	target = 1

	result = combinationSum2(candidates, target)
	if len(result) != 1 {
		t.Errorf("test candidates: %+v, target: %d failed", candidates, target)
		t.Errorf("got result: %+v, ", result)
	}
}