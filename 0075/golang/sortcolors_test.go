package sortcolors

import (
	"testing"
)

func TestSortColors(t *testing.T) {
	input := []int{2, 0, 2, 1, 1, 0}
	target := []int{0, 0, 1, 1, 2, 2}

	sortColors(input)
	if !isSame(input, target) {
		t.Errorf("expect %+v, got: %+v", target, input)
	}

	input = []int{2, 2, 1, 1, 0, 0}
	target = []int{0, 0, 1, 1, 2, 2}

	sortColors(input)
	if !isSame(input, target) {
		t.Errorf("expect %+v, got: %+v", target, input)
	}

	input = []int{0, 0, 1, 1, 2, 2}
	target = []int{0, 0, 1, 1, 2, 2}

	sortColors(input)
	if !isSame(input, target) {
		t.Errorf("expect %+v, got: %+v", target, input)
	}

	input = []int{1, 0, 2}
	target = []int{0, 1, 2}

	sortColors(input)
	if !isSame(input, target) {
		t.Errorf("expect %+v, got: %+v", target, input)
	}

	input = []int{1, 2, 0}
	target = []int{0, 1, 2}

	sortColors(input)
	if !isSame(input, target) {
		t.Errorf("expect %+v, got: %+v", target, input)
	}

}

func isSame(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}

	same := true
	for idx, num1 := range nums1 {
		if num1 != nums2[idx] {
			same = false
			break
		}
	}
	return same
}
