package besttimetobuyandsellstockiv

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	input := []int{2, 4, 1}
	k := 2
	target := 2

	result := maxProfit(k, input)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMaxProfitII(t *testing.T) {
	input := []int{3, 2, 6, 5, 0, 3}
	k := 2
	target := 7

	result := maxProfit(k, input)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMaxProfitIII(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	target := 6

	result := maxProfit(k, input)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMaxProfitIV(t *testing.T) {
	input := []int{7, 6, 5, 4, 3, 2, 1}
	k := 3
	target := 0

	result := maxProfit(k, input)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMaxProfitV(t *testing.T) {
	input := []int{6, 1, 3, 2, 4, 7}
	k := 2
	target := 7

	result := maxProfit(k, input)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMaxProfitVI(t *testing.T) {
	input := []int{3, 3, 5, 0, 0, 3, 1, 4}
	k := 2
	target := 6

	result := maxProfit(k, input)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMaxProfitVII(t *testing.T) {
	input := []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}
	k := 2
	target := 13

	result := maxProfit(k, input)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
