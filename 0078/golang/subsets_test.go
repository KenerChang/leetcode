package subsets

import (
	"fmt"
	"testing"
)

func TestSubsets(t *testing.T) {
	input := []int{1, 2, 3}
	target := [][]int{
		[]int{},
		[]int{1}, []int{2}, []int{3},
		[]int{1, 2}, []int{2, 3}, []int{1, 3},
		[]int{1, 2, 3},
	}

	output := subsets(input)
	if !isSame(target, output) {
		t.Errorf("expect %+v, got: %+v", target, output)
	}
}

func isSame(set1, set2 [][]int) bool {
	map1 := map[string]bool{}

	if len(set1) != len(set2) {
		return false
	}

	for _, nums := range set1 {
		key := fmt.Sprint(nums)
		map1[key] = true
	}

	same := true
	for _, nums := range set2 {
		key := fmt.Sprint(nums)

		if _, ok := map1[key]; !ok {
			same = false
			break
		}
	}
	return same
}
