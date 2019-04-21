package searchmatrix

import (
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	input := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	searchTarget := 3
	target := true

	result := searchMatrix(input, searchTarget)
	if result != target {
		t.Errorf("expect %t, got: %t", target, result)
	}

	input = [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	searchTarget = 13
	target = false

	result = searchMatrix(input, searchTarget)
	if result != target {
		t.Errorf("expect %t, got: %t", target, result)
	}

	input = [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	searchTarget = 10
	target = true

	result = searchMatrix(input, searchTarget)
	if result != target {
		t.Errorf("expect %t, got: %t", target, result)
	}

	input = [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	searchTarget = 50
	target = true

	result = searchMatrix(input, searchTarget)
	if result != target {
		t.Errorf("expect %t, got: %t", target, result)
	}

	input = [][]int{
		[]int{1},
	}
	searchTarget = 1
	target = true

	result = searchMatrix(input, searchTarget)
	if result != target {
		t.Errorf("expect %t, got: %t", target, result)
	}

	input = [][]int{
		[]int{1},
	}
	searchTarget = 2
	target = false

	result = searchMatrix(input, searchTarget)
	if result != target {
		t.Errorf("expect %t, got: %t", target, result)
	}
}
