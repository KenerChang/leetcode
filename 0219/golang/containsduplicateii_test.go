package containsduplicateii

import (
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	k := 3

	result := containsNearbyDuplicate(nums, k)
	if !result {
		t.Errorf("expect ture, but got false")
	}
}

func TestContainsNearbyDuplicateII(t *testing.T) {
	nums := []int{1, 0, 1, 1}
	k := 1

	result := containsNearbyDuplicate(nums, k)
	if !result {
		t.Errorf("expect ture, but got false")
	}
}

func TestContainsNearbyDuplicateIII(t *testing.T) {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 1

	result := containsNearbyDuplicate(nums, k)
	if result {
		t.Errorf("expect false, but got true")
	}
}

func TestContainsNearbyDuplicateIV(t *testing.T) {
	nums := []int{1, 1}
	k := 0

	result := containsNearbyDuplicate(nums, k)
	if result {
		t.Errorf("expect false, but got true")
	}
}

func TestContainsNearbyDuplicateV(t *testing.T) {
	nums := []int{1}
	k := 10

	result := containsNearbyDuplicate(nums, k)
	if result {
		t.Errorf("expect false, but got true")
	}
}
