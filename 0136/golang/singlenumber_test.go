package singlenumber

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	nums := []int{2, 2, 1}
	target := 1

	result := singleNumber(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestSingleNumberII(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2}
	target := 4

	result := singleNumber(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestSingleNumberIII(t *testing.T) {
	nums := []int{4, 4, 0}
	target := 0

	result := singleNumber(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
