package rotatearray

import (
	"testing"
)

func verify(target, result []int) bool {
	if len(target) != len(result) {
		return false
	}

	for i, nt := range target {
		if nt != result[i] {
			return false
		}
	}
	return true
}

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	target := []int{5, 6, 7, 1, 2, 3, 4}

	rotate(nums, k)
	if !verify(nums, target) {
		t.Errorf("expect %v, got %v", target, nums)
	}

}

func TestRotateII(t *testing.T) {
	nums := []int{-1, -100, 3, 99}
	k := 2
	target := []int{3, 99, -1, -100}

	rotate(nums, k)
	if !verify(nums, target) {
		t.Errorf("expect %v, got %v", target, nums)
	}
}

func TestRotateIII(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6}
	k := 0
	target := []int{0, 1, 2, 3, 4, 5, 6}

	rotate(nums, k)
	if !verify(nums, target) {
		t.Errorf("expect %v, got %v", target, nums)
	}
}

func TestRotateIV(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6}
	k := 1
	target := []int{6, 0, 1, 2, 3, 4, 5}

	rotate(nums, k)
	if !verify(nums, target) {
		t.Errorf("expect %v, got %v", target, nums)
	}
}

func TestRotateV(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6}
	k := 6
	target := []int{1, 2, 3, 4, 5, 6, 0}

	rotate(nums, k)
	if !verify(nums, target) {
		t.Errorf("expect %v, got %v", target, nums)
	}
}

func TestRotateVI(t *testing.T) {
	nums := []int{1, 2}
	k := 3
	target := []int{2, 1}

	rotate(nums, k)
	if !verify(nums, target) {
		t.Errorf("expect %v, got %v", target, nums)
	}
}

func TestRotateVII(t *testing.T) {
	nums := []int{1, 2}
	k := 4
	target := []int{1, 2}

	rotate(nums, k)
	if !verify(nums, target) {
		t.Errorf("expect %v, got %v", target, nums)
	}
}

func TestRotateVIII(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	k := 4
	target := []int{3, 4, 5, 6, 1, 2}

	rotate(nums, k)
	if !verify(nums, target) {
		t.Errorf("expect %v, got %v", target, nums)
	}
}

func TestRotateIX(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	k := 5
	target := []int{2, 3, 4, 5, 6, 1}

	rotate(nums, k)
	if !verify(nums, target) {
		t.Errorf("expect %v, got %v", target, nums)
	}
}

func TestRotateX(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}
	k := 38
	target := []int{17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	rotate(nums, k)
	if !verify(nums, target) {
		t.Errorf("expect %v, got %v", target, nums)
	}
}
