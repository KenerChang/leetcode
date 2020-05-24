package triangle

import (
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	tri := [][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	}
	target := 11

	result := minimumTotal(tri)

	if target != result {
		t.Errorf("expect: %d, got: %d", target, result)
	}
}

func TestMinimumTotalII(t *testing.T) {
	tri := [][]int{
		[]int{2},
	}
	target := 2

	result := minimumTotal(tri)

	if target != result {
		t.Errorf("expect: %d, got: %d", target, result)
	}
}

func TestMinimumTotalIII(t *testing.T) {
	tri := [][]int{
		[]int{2},
		[]int{3, 4},
	}
	target := 5

	result := minimumTotal(tri)

	if target != result {
		t.Errorf("expect: %d, got: %d", target, result)
	}
}

func TestMinimumTotalIV(t *testing.T) {
	tri := [][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, -5, 7},
		[]int{4, 10, 8, 3},
	}
	target := 8

	result := minimumTotal(tri)

	if target != result {
		t.Errorf("expect: %d, got: %d", target, result)
	}
}
