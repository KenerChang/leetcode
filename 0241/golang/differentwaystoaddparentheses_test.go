package differentwaystoaddparentheses

import (
	"sort"
	"testing"
)

func verify(target, result []int) bool {
	if len(target) != len(result) {
		return false
	}

	sort.Ints(target)
	sort.Ints(result)

	for i := 0; i < len(target); i++ {
		if target[i] != result[i] {
			return false
		}
	}
	return true
}

func TestDiffWaysToCompute(t *testing.T) {
	result := diffWaysToCompute("2-1-1")

	target := []int{0, 2}

	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestDiffWaysToComputeII(t *testing.T) {
	result := diffWaysToCompute("2*3-4*5")

	target := []int{-34, -14, -10, -10, 10}

	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}
