package kthlargestelementinanarray

import (
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	target := 5

	result := findKthLargest(nums, k)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestFindKthLargestII(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	target := 4

	result := findKthLargest(nums, k)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestFindKthLargestIII(t *testing.T) {
	nums := []int{}
	k := 1
	target := -1

	result := findKthLargest(nums, k)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestFindKthLargestIV(t *testing.T) {
	nums := []int{1, 1, 1, 1}
	k := 2
	target := 1

	result := findKthLargest(nums, k)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestFindKthLargestV(t *testing.T) {
	nums := []int{-1, 2, 0}
	k := 2
	target := 0

	result := findKthLargest(nums, k)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
