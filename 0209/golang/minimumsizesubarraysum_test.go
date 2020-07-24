package minimumsizesubarraysum

import (
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	s := 7
	target := 2

	result := minSubArrayLen(s, nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMinSubArrayLenII(t *testing.T) {
	nums := []int{}
	s := 7
	target := 0

	result := minSubArrayLen(s, nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMinSubArrayLenIII(t *testing.T) {
	nums := []int{1, 1, 1}
	s := 7
	target := 0

	result := minSubArrayLen(s, nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMinSubArrayLenIV(t *testing.T) {
	nums := []int{7, 2, 5, 7}
	s := 7
	target := 1

	result := minSubArrayLen(s, nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMinSubArrayLenV(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	s := 11
	target := 3

	result := minSubArrayLen(s, nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
