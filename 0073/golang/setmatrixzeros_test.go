package setmatrixzeros

import (
	"testing"
)

func TestSetZeros(t *testing.T) {
	input := [][]int{
		[]int{1, 1, 1},
		[]int{1, 0, 1},
		[]int{1, 1, 1},
	}
	target := [][]int{
		[]int{1, 0, 1},
		[]int{0, 0, 0},
		[]int{1, 0, 1},
	}

	setZeroes(input)
	if !isSame(input, target) {
		t.Errorf("expect %+v, got %+v", target, input)
	}

	input = [][]int{
		[]int{0, 1, 2, 0},
		[]int{3, 4, 5, 2},
		[]int{1, 3, 1, 5},
	}
	target = [][]int{
		[]int{0, 0, 0, 0},
		[]int{0, 4, 5, 0},
		[]int{0, 3, 1, 0},
	}

	setZeroes(input)
	if !isSame(input, target) {
		t.Errorf("expect %+v, got %+v", target, input)
	}
}

func isSame(m1, m2 [][]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for ri, row := range m1 {
		for ci, col := range row {
			if col != m2[ri][ci] {
				return false
			}
		}
	}
	return true
}
