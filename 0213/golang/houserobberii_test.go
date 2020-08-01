package houserobberii

import (
	"testing"
)

func TestRob(t *testing.T) {
	nums := []int{2, 3, 2}
	target := 3

	result := rob(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestRobII(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	target := 4

	result := rob(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestRobIII(t *testing.T) {
	nums := []int{1, 2, 3, 1, 1, 4}
	target := 7

	result := rob(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestRobIV(t *testing.T) {
	nums := []int{1, 1, 1, 1}
	target := 2

	result := rob(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestRobV(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 30

	result := rob(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestRobVI(t *testing.T) {
	nums := []int{94, 40, 49, 65, 21, 21, 106, 80, 92, 81, 679, 4, 61, 6, 237, 12, 72, 74, 29, 95, 265, 35, 47, 1, 61, 397, 52, 72, 37, 51, 1, 81, 45, 435, 7, 36, 57, 86, 81, 72}
	target := 2926

	result := rob(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
