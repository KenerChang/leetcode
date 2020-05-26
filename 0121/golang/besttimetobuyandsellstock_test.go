package besttimetobuyandsellstock

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	target := 5

	result := maxProfit(prices)
	if target != result {
		t.Errorf("expect %d, but got %d", target, result)
	}
}

func TestMaxProfitII(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	target := 0

	result := maxProfit(prices)
	if target != result {
		t.Errorf("expect %d, but got %d", target, result)
	}
}

func TestMaxProfitIII(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5}
	target := 4

	result := maxProfit(prices)
	if target != result {
		t.Errorf("expect %d, but got %d", target, result)
	}
}

func TestMaxProfitIV(t *testing.T) {
	prices := []int{1}
	target := 0

	result := maxProfit(prices)
	if target != result {
		t.Errorf("expect %d, but got %d", target, result)
	}
}
