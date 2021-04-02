package containsduplicateiii

import "testing"

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	result := containsNearbyAlmostDuplicate(nums, 3, 0)

	if !result {
		t.Errorf("expect true, got false")
	}
}

func TestContainsNearbyAlmostDuplicateII(t *testing.T) {
	nums := []int{1, 0, 1, 1}
	result := containsNearbyAlmostDuplicate(nums, 1, 2)

	if !result {
		t.Errorf("expect true, got false")
	}
}

func TestContainsNearbyAlmostDuplicateIII(t *testing.T) {
	nums := []int{1, 5, 9, 1, 5, 9}
	result := containsNearbyAlmostDuplicate(nums, 2, 3)

	if result {
		t.Errorf("expect false, got true")
	}
}

func TestContainsNearbyAlmostDuplicateIV(t *testing.T) {
	nums := []int{1}
	result := containsNearbyAlmostDuplicate(nums, 0, 0)

	if result {
		t.Errorf("expect false, got true")
	}
}
