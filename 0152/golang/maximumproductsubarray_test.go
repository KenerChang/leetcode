package maximumproductsubarray

import (
	"testing"
)

func TestMaxProduct(t *testing.T) {
	nums := []int{2, 3, -2, 4}
	target := 6

	result := maxProduct(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMaxProductII(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	target := 24

	result := maxProduct(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMaxProductIII(t *testing.T) {
	nums := []int{-1, -2, -3, -4}
	target := 24

	result := maxProduct(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMaxProductIV(t *testing.T) {
	nums := []int{5}
	target := 5

	result := maxProduct(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMaxProductV(t *testing.T) {
	nums := []int{2, -2, -3, -4}
	target := 12

	result := maxProduct(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMaxProductVI(t *testing.T) {
	nums := []int{}
	target := 0

	result := maxProduct(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMaxProductVII(t *testing.T) {
	nums := []int{0, 5, -100, -2, -3, 10}
	target := 1000

	result := maxProduct(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMaxProductVIII(t *testing.T) {
	nums := []int{0, 2}
	target := 2

	result := maxProduct(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMaxProductVIIII(t *testing.T) {
	nums := []int{-3, 0, 1, -2}
	target := 1

	result := maxProduct(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMaxProductX(t *testing.T) {
	nums := []int{5, -100, -2, -3, 10}
	target := 1000

	result := maxProduct(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
