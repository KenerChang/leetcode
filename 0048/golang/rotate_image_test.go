package rotate

import (
	"testing"
)

func TestRotate(t *testing.T) {
    matrix := [][]int {
		[]int{1,2,3},
		[]int{4,5,6},
		[]int{7,8,9},
	}
	target := [][]int {
		[]int{7,4,1},
		[]int{8,5,2},
		[]int{9,6,3},
	}

	rotate(matrix)
	if !isEqual(matrix, target) {
		t.Errorf("expect %+v, got: %+v", target, matrix)
	}


	matrix = [][]int {
		[]int{1,2,3,4},
		[]int{5,6,7,8},
		[]int{9,10,11,12},
		[]int{13,14,15,16},
	}
	target = [][]int {
		[]int{13,9,5,1},
		[]int{14,10,6,2},
		[]int{15,11,7,3},
		[]int{16,12,8,4},
	}

	rotate(matrix)
	if !isEqual(matrix, target) {
		t.Errorf("expect %+v, got: %+v", target, matrix)
	}
}

func isEqual(m1, m2 [][]int) bool {
	equal := true
	for i, row := range m1 {
		for j, col := range row {
			if col != m2[i][j] {
				equal = false
				break
			}
		}
	}
	return equal
}