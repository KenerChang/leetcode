package besttimetobuyandsellstockii

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	target := 7

	result := maxProfit(prices)
	if target != result {
		t.Errorf("expect %d, but got %d", target, result)
	}
}

func TestMaxProfitII(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5}
	target := 4

	result := maxProfit(prices)
	if target != result {
		t.Errorf("expect %d, but got %d", target, result)
	}
}

func TestMaxProfitIII(t *testing.T) {
	prices := []int{5, 4, 3, 2, 1}
	target := 0

	result := maxProfit(prices)
	if target != result {
		t.Errorf("expect %d, but got %d", target, result)
	}
}

func TestMaxProfitIV(t *testing.T) {
	prices := []int{1, 2, 2, 3, 3, 4, 4, 5, 5, 5}
	target := 4

	result := maxProfit(prices)
	if target != result {
		t.Errorf("expect %d, but got %d", target, result)
	}
}

func TestMaxProfitV(t *testing.T) {
	prices := []int{7, 1, 5, 5, 5, 3, 6, 4}
	target := 7

	result := maxProfit(prices)
	if target != result {
		t.Errorf("expect %d, but got %d", target, result)
	}
}

func TestMaxProfitVI(t *testing.T) {
	prices := []int{7, 1, 5, 5, 5, 3, 6, 4, 1, 1, 6}
	target := 12

	result := maxProfit(prices)
	if target != result {
		t.Errorf("expect %d, but got %d", target, result)
	}
}
