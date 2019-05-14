package mergesortedarray

import (
	"testing"
)

func TestMerge(t *testing.T) {
	template := "expect %+v, got %+v"
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	target := []int{1, 2, 2, 3, 5, 6}

	merge(nums1, 3, nums2, 3)

	if !isSame(nums1, target) {
		t.Errorf(template, nums1, nums2)
	}

	nums1 = []int{1, 2, 3, 4, 0, 0, 0}
	nums2 = []int{2, 5, 6}
	target = []int{1, 2, 2, 3, 4, 5, 6}

	merge(nums1, 4, nums2, 3)

	if !isSame(nums1, target) {
		t.Errorf(template, nums1, nums2)
	}

	nums1 = []int{1, 2, 3, 4, 0}
	nums2 = []int{2}
	target = []int{1, 2, 2, 3, 4}

	merge(nums1, 4, nums2, 1)

	if !isSame(nums1, target) {
		t.Errorf(template, target, nums1)
	}

	nums1 = []int{1, 2, 3, 4}
	nums2 = []int{}
	target = []int{1, 2, 3, 4}

	merge(nums1, 4, nums2, 0)

	if !isSame(nums1, target) {
		t.Errorf(template, target, nums1)
	}

	nums1 = []int{0, 0}
	nums2 = []int{1}
	target = []int{0, 1}

	merge(nums1, 1, nums2, 1)

	if !isSame(nums1, target) {
		t.Errorf(template, target, nums1)
	}

}

func isSame(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}

	for ind, num := range nums1 {
		if num != nums2[ind] {
			return false
		}
	}
	return true
}
