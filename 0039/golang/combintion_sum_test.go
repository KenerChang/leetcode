package combination_sum

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	// test candidates = [2,3,6,7], target = 7
	candidates := []int{2,3,6,7}
	target := 7

	result := combinationSum(candidates, target)
	fmt.Printf("result: %+v\n", result)
	if len(result) != 2 {
		t.Errorf("test candidates: %+v, target: %d failed", candidates, target)
	}

	// test candidates = [2,3,5], target = 8
	candidates = []int{2,3,5}
	target = 8
	result = combinationSum(candidates, target)
	fmt.Printf("result: %+v\n", result)
	if len(result) != 3 {
		t.Errorf("test candidates: %+v, target: %d failed", candidates, target)
	}

	// test candidates = [0,1,2,3] target = 0
	candidates = []int{0,1,2,3}
	target = 0
	result = combinationSum(candidates, target)
	fmt.Printf("result: %+v\n", result)
	if len(result) != 0 {
		t.Errorf("test candidates: %+v, target: %d failed", candidates, target)
	}

	// test candidates = [7,3,2] target = 18
	candidates = []int{7,3,2}
	target = 18
	result = combinationSum(candidates, target)
	fmt.Printf("result: %+v\n", result)
	if len(result) != 7 {
		t.Errorf("test candidates: %+v, target: %d failed", candidates, target)
	}
}