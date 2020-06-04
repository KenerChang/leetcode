package gasstation

import (
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}

	target := 3
	result := canCompleteCircuit(gas, cost)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCanCompleteCircuitII(t *testing.T) {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}

	target := -1
	result := canCompleteCircuit(gas, cost)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCanCompleteCircuitIII(t *testing.T) {
	gas := []int{-2}
	cost := []int{-2}

	target := 0
	result := canCompleteCircuit(gas, cost)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCanCompleteCircuitIV(t *testing.T) {
	gas := []int{100, 5}
	cost := []int{100, 5}

	target := 0
	result := canCompleteCircuit(gas, cost)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCanCompleteCircuitV(t *testing.T) {
	gas := []int{}
	cost := []int{}

	target := -1
	result := canCompleteCircuit(gas, cost)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
