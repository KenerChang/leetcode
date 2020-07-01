package dungeongame

import (
	"testing"
)

func TestCalCulateMinimumHP(t *testing.T) {
	dungeon := [][]int{
		[]int{-2, -3, 3},
		[]int{-5, -10, 1},
		[]int{10, 30, -5},
	}

	target := 7
	result := calculateMinimumHP(dungeon)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCalCulateMinimumHPII(t *testing.T) {
	dungeon := [][]int{
		[]int{-2},
	}

	target := 3
	result := calculateMinimumHP(dungeon)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCalCulateMinimumHPIII(t *testing.T) {
	dungeon := [][]int{
		[]int{2, 3, 3},
		[]int{5, 10, 1},
		[]int{10, 30, 5},
	}

	target := 1
	result := calculateMinimumHP(dungeon)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCalCulateMinimumHPIV(t *testing.T) {
	dungeon := [][]int{
		[]int{-2, -3, -3},
		[]int{-5, -10, -1},
		[]int{-10, -30, -5},
	}

	target := 15
	result := calculateMinimumHP(dungeon)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
