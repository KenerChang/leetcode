package miniumpathsum

import (
	"testing"
)

func TestMinPathSum(t *testing.T) {
	inputs := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}
	target := 7

	result := minPathSum(inputs)
	if result != target {
		t.Errorf("expect %d, but got %d", target, result)
	}

	inputs = [][]int{
		[]int{1},
	}
	target = 1

	result = minPathSum(inputs)
	if result != target {
		t.Errorf("expect %d, but got %d", target, result)
	}

	inputs = [][]int{
		[]int{},
	}
	target = 0

	result = minPathSum(inputs)
	if result != target {
		t.Errorf("expect %d, but got %d", target, result)
	}

	inputs = [][]int{
		[]int{1},
		[]int{1},
	}
	target = 2

	result = minPathSum(inputs)
	if result != target {
		t.Errorf("expect %d, but got %d", target, result)
	}
}
