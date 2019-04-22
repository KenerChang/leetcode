package combinations

import (
	"testing"
)

func TestCombine(t *testing.T) {
	n := 4
	k := 2
	target := [][]int{
		[]int{2, 4},
		[]int{3, 4},
		[]int{2, 3},
		[]int{1, 2},
		[]int{1, 3},
		[]int{1, 4},
	}

	result := combine(n, k)
	if len(result) != len(target) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}
