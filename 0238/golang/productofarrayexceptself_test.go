package productofarrayexceptself

import "testing"

func verify(target, result []int) bool {
	if len(target) != len(result) {
		return false
	}

	for i := range target {
		if target[i] != result[i] {
			return false
		}
	}
	return true
}

func TestProductExceptSelf(t *testing.T) {
	nums := []int{1, 2, 3, 4}

	result := productExceptSelf(nums)

	target := []int{24, 12, 8, 6}
	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestProductExceptSelfII(t *testing.T) {
	nums := []int{-1, 1, 0, -3, 3}

	result := productExceptSelf(nums)

	target := []int{0, 0, 9, 0, 0}
	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}
