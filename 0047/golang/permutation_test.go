package permutation

import (
	"testing"
	"fmt"
)

func TestPermuteUnique(t *testing.T) {
	input := []int{
		1,1,2,
	}
	target := [][]int {
		[]int{
			1,1,2,
		},
		[]int {
			1,2,1,
		},
		[]int {
			2,1,1,
		},
	}
	result := permuteUnique(input)
	if !isEqual(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}

	input = []int {
		3,3,3,
	}
	target = [][]int {
		[]int{
			3,3,3,
		},
	} 
	result = permuteUnique(input)
	if !isEqual(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}

	input = []int {}
	target = [][]int {} 
	result = permuteUnique(input)
	if !isEqual(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}

func isEqual(nums1, nums2 [][]int) bool {
	nums1Map := map[string]bool{}
	for _, nums := range nums1 {
		key := fmt.Sprint(nums)
		nums1Map[key] = true
	}

	nums2Map := map[string]bool{}
	for _, nums := range nums2 {
		key := fmt.Sprint(nums)
		nums2Map[key] = true
	}

	if len(nums1Map) != len(nums2Map) {
		return false
	}

	equal := true
	for k := range nums1Map {
		if _, ok := nums2Map[k]; !ok {
			equal = false
			break
		}
	}
	return equal
}