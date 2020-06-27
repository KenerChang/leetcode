package twosumiiinputarrayissorted

import (
	"testing"
)

func verify(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}

	for i, num1 := range nums1 {
		if num1 != nums2[i] {
			return false
		}
	}

	return true
}

func TestTwoSum(t *testing.T) {
	target := []int{1, 2}
	result := twoSum([]int{2, 7, 11, 15}, 9)

	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}
