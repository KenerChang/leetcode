package maximumsubarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	result := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})

	target := 6
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMaxSubArrayII(t *testing.T) {
	result := maxSubArray([]int{5, 4, -1, 7, 8})

	target := 23
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMaxSubArrayIII(t *testing.T) {
	result := maxSubArray([]int{-1, -2})

	target := -1
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}
