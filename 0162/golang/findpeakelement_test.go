package findpeakelement

import (
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	target := 2

	result := findPeakElement(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestFindPeakElementII(t *testing.T) {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	target := 1

	result := findPeakElement(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestFindPeakElementIII(t *testing.T) {
	nums := []int{1, 2, 3}
	target := 2

	result := findPeakElement(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestFindPeakElementIV(t *testing.T) {
	nums := []int{3, 2, 1}
	target := 0

	result := findPeakElement(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
