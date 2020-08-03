package combinationsumiii

import (
	"testing"
)

func verify(target, result [][]int) bool {
	if len(target) != len(result) {
		return false
	}

	for i, numst := range target {
		if len(numst) != len(result[i]) {
			return false
		}

		for j, numt := range numst {
			if numt != result[i][j] {
				return false
			}
		}
	}
	return true
}

func TestCombinationSum3(t *testing.T) {
	target := [][]int{
		[]int{1, 2, 4},
	}

	result := combinationSum3(3, 7)
	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestCombinationSum3II(t *testing.T) {
	target := [][]int{
		[]int{1, 2, 6},
		[]int{1, 3, 5},
		[]int{2, 3, 4},
	}

	result := combinationSum3(3, 9)
	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestCombinationSum3III(t *testing.T) {
	target := [][]int{
		[]int{8, 9},
	}

	result := combinationSum3(2, 17)
	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestCombinationSum3IV(t *testing.T) {
	target := [][]int{}

	result := combinationSum3(0, 1)
	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestCombinationSum3V(t *testing.T) {
	target := [][]int{}

	result := combinationSum3(1, 10)
	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestCombinationSum3VI(t *testing.T) {
	target := [][]int{}

	result := combinationSum3(9, 46)
	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestCombinationSum3VII(t *testing.T) {
	target := [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	result := combinationSum3(9, 45)
	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}
