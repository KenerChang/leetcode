package pascalstriangle

import (
	"testing"
)

func verify(target, result [][]int) bool {
	if len(target) != len(result) {
		return false
	}

	for i, nums1 := range target {
		nums2 := result[i]
		if len(nums1) != len(nums2) {
			return false
		}

		for j, num1 := range nums1 {
			num2 := nums2[j]
			if num1 != num2 {
				return false
			}
		}
	}
	return true
}

func TestGenerate(t *testing.T) {
	target := [][]int{
		[]int{1},
		[]int{1, 1},
		[]int{1, 2, 1},
		[]int{1, 3, 3, 1},
		[]int{1, 4, 6, 4, 1},
	}
	result := generate(5)

	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}

func TestGenerateII(t *testing.T) {
	target := [][]int{
		[]int{1},
	}
	result := generate(1)

	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}

func TestGenerateIII(t *testing.T) {
	target := [][]int{
		[]int{1},
		[]int{1, 1},
		[]int{1, 2, 1},
	}
	result := generate(3)

	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}

func TestGenerateIV(t *testing.T) {
	target := [][]int{
		[]int{1},
		[]int{1, 1},
	}
	result := generate(2)

	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}

func TestGenerateV(t *testing.T) {
	target := [][]int{
		[]int{1},
		[]int{1, 1},
		[]int{1, 2, 1},
		[]int{1, 3, 3, 1},
		[]int{1, 4, 6, 4, 1},
		[]int{1, 5, 10, 10, 5, 1},
	}
	result := generate(6)

	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}
