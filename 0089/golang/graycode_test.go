package graycode

import (
	"testing"
)

func TestGrayCode(t *testing.T) {
	template := "expect %+v, got: %+v"
	input := 2
	target := []int{0, 1, 3, 2}

	result := grayCode(input)
	if !isSame(target, result) {
		t.Errorf(template, target, result)
	}

	input = 1
	target = []int{0, 1}

	result = grayCode(input)
	if !isSame(target, result) {
		t.Errorf(template, target, result)
	}

	input = 0
	target = []int{0}

	result = grayCode(input)
	if !isSame(target, result) {
		t.Errorf(template, target, result)
	}

	input = 3
	target = []int{0, 1, 3, 2, 6, 7, 5, 4}

	result = grayCode(input)
	if !isSame(target, result) {
		t.Errorf(template, target, result)
	}

	input = 4
	target = []int{0, 1, 3, 2, 6, 7, 5, 4, 12, 13, 15, 14, 10, 11, 9, 8}

	result = grayCode(input)
	if !isSame(target, result) {
		t.Errorf(template, target, result)
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
