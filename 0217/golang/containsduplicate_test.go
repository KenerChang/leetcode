package containsduplicate

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	target := true

	result := containsDuplicate(nums)
	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestContainsDuplicateII(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	target := false

	result := containsDuplicate(nums)
	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestContainsDuplicateIII(t *testing.T) {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	target := true

	result := containsDuplicate(nums)
	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}
