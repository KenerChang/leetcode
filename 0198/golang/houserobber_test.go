package houserobber

import (
	"testing"
)

func TestRob(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	target := 4

	result := rob(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestRobII(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	target := 12

	result := rob(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestRobIII(t *testing.T) {
	nums := []int{}
	target := 0

	result := rob(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestRobIV(t *testing.T) {
	nums := []int{1}
	target := 1

	result := rob(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestRobV(t *testing.T) {
	nums := []int{1, 2}
	target := 2

	result := rob(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestRobVI(t *testing.T) {
	nums := []int{1, 2, 3}
	target := 4

	result := rob(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestRobVII(t *testing.T) {
	nums := []int{1, 2, 3, 2}
	target := 4

	result := rob(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
