package singlenumberii

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	nums := []int{2, 2, 3, 2}
	target := 3

	result := singleNumber(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestSingleNumberII(t *testing.T) {
	nums := []int{0, 1, 0, 1, 0, 1, 99}
	target := 99

	result := singleNumber(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
