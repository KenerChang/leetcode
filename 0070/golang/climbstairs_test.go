package climbstairs

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	input := 2
	target := 2

	result := climbStairs(input)
	if result != target {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	input = 3
	target = 3

	result = climbStairs(input)
	if result != target {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	input = 0
	target = 1

	result = climbStairs(input)
	if result != target {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	input = 10
	target = 89

	result = climbStairs(input)
	if result != target {
		t.Errorf("expect %d, but got: %d", target, result)
	}
}
