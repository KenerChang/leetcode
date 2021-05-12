package painthouse

import "testing"

func TestMinCost(t *testing.T) {
	costs := [][]int{
		{17, 2, 17}, {16, 16, 5}, {14, 3, 19},
	}

	result := minCost(costs)

	target := 10
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMinCostII(t *testing.T) {
	costs := [][]int{
		{7, 6, 2},
	}

	result := minCost(costs)

	target := 2
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
