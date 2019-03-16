package permutations

import (
	"testing"
)

func TestPermute(t *testing.T) {
	input := []int{1,2,3}
	target := [][]int{
		[]int{
			1,2,3,
		},
		[]int{
			1,3,2,
		},
		[]int{
			2,1,3,
		},
		[]int{
			2,3,1,
		},
		[]int{
			3,1,2,
		},
		[]int{
			3,2,1,
		},
	}

	result := permute(input)
	if !isEqual(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}

	input = []int{1,2,3,4}
	target = [][]int{
		[]int{
			1,2,3,4,
		},
		[]int{
			1,2,4,3,
		},
		[]int {
			1,3,2,4,
		},
		[]int {
			1,3,4,2,
		},
		[]int {
			1,4,2,3,
		},
		[]int {
			1,4,3,2,
		},
		[]int {
			2,1,3,4,
		},
		[]int{
			2,1,4,3,
		},
		[]int{
			2,3,1,4,
		},
		[]int {
			2,3,4,1,
		},
		[]int {
			2,4,1,3,
		},
		[]int {
			2,4,3,1,
		},
		[]int {
			3,1,2,4,
		},
		[]int{
			3,1,4,2,
		},
		[]int {
			3,2,1,4,
		},
		[]int {
			3,2,4,1,
		},
		[]int {
			3,4,1,2,
		},
		[]int {
			3,4,2,1,
		},
		[]int{
			4,1,2,3,
		},
		[]int {
			4,1,3,2,
		},
		[]int {
			4,2,1,3,
		},
		[]int {
			4,2,3,1,
		},
		[]int {
			4,3,1,2,
		},
		[]int {
			4,3,2,1,
		},
	}
	result = permute(input)
	if !isEqual(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}

	input = []int{}
	target = [][]int{}
	result = permute(input)
	if !isEqual(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}

	input = []int{1}
	target = [][]int{
		[]int{1},
	}
	result = permute(input)
	if !isEqual(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}

func isEqual(nums1, nums2 [][]int) bool {
	if len(nums1) != len(nums2) {
		return false
	}

	equal := true
	for i, nums := range nums1 {
		for j, num := range nums {
			if num != nums2[i][j] {
				equal = false
				break
			}
		}
		if !equal {
			break
		}
	}
	return equal
}