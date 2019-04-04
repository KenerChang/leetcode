package spiralmatrixII

import (
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	target := [][]int{
		[]int{1, 2, 3},
		[]int{8, 9, 4},
		[]int{7, 6, 5},
	}

	result := generateMatrix(3)
	if !isSame(target, result) {
		t.Errorf("expect %+v, but got: %+v", target, result)
	}

	target = [][]int{
		[]int{1, 2, 3, 4},
		[]int{12, 13, 14, 5},
		[]int{11, 16, 15, 6},
		[]int{10, 9, 8, 7},
	}

	result = generateMatrix(4)
	if !isSame(target, result) {
		t.Errorf("expect %+v, but got: %+v", target, result)
	}

	target = [][]int{}

	result = generateMatrix(0)
	if !isSame(target, result) {
		t.Errorf("expect %+v, but got: %+v", target, result)
	}

	target = [][]int{
		[]int{1},
	}

	result = generateMatrix(1)
	if !isSame(target, result) {
		t.Errorf("expect %+v, but got: %+v", target, result)
	}
}

func isSame(m1, m2 [][]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	result := true
	for i, row1 := range m1 {
		if len(row1) != len(m2[i]) {
			result = false
			break
		}

		for j, col1 := range row1 {
			col2 := m2[i][j]
			if col1 != col2 {
				result = false
				break
			}
		}
	}
	return result
}
