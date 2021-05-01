package singlenumberiii

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

func TestSingleNumber(t *testing.T) {
	nums := []int{1, 2, 1, 3, 2, 5}

	result := singleNumber(nums)

	target := []int{3, 5}
	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestSingleNumberII(t *testing.T) {
	nums := []int{3, 5}

	result := singleNumber(nums)

	target := []int{3, 5}
	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}
