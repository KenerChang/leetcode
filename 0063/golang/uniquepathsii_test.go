package uniquepathsii

import (
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	matrix := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}
	target := 2

	result := uniquePathsWithObstacles(matrix)
	if result != target {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	matrix = [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 1},
		[]int{0, 0, 0},
	}
	target = 1

	result = uniquePathsWithObstacles(matrix)
	if result != target {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	matrix = [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 1},
		[]int{0, 0, 0},
	}
	target = 1

	result = uniquePathsWithObstacles(matrix)
	if result != target {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	matrix = [][]int{
		[]int{0},
	}
	target = 1

	result = uniquePathsWithObstacles(matrix)
	if result != target {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	matrix = [][]int{
		[]int{1},
	}
	target = 0

	result = uniquePathsWithObstacles(matrix)
	if result != target {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	matrix = [][]int{
		[]int{0, 1},
	}
	target = 0

	result = uniquePathsWithObstacles(matrix)
	if result != target {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	matrix = [][]int{
		[]int{0},
		[]int{1},
	}
	target = 0

	result = uniquePathsWithObstacles(matrix)
	if result != target {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	matrix = [][]int{
		[]int{1},
		[]int{0},
	}
	target = 0

	result = uniquePathsWithObstacles(matrix)
	if result != target {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	matrix = [][]int{
		[]int{1, 0},
	}
	target = 0

	result = uniquePathsWithObstacles(matrix)
	if result != target {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	matrix = [][]int{
		[]int{0, 1},
		[]int{1, 0},
	}
	target = 0

	result = uniquePathsWithObstacles(matrix)
	if result != target {
		t.Errorf("expect %d, but got: %d", target, result)
	}
}
