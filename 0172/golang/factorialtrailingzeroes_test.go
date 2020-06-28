package factorialtrailingzeroes

import (
	"testing"
)

func TestTrailingZeroes(t *testing.T) {
	target := 0
	result := trailingZeroes(3)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestTrailingZeroesII(t *testing.T) {
	target := 1
	result := trailingZeroes(5)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestTrailingZeroesIII(t *testing.T) {
	target := 2
	result := trailingZeroes(10)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestTrailingZeroesIV(t *testing.T) {
	target := 6
	result := trailingZeroes(25)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestTrailingZeroesV(t *testing.T) {
	target := 24
	result := trailingZeroes(100)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestTrailingZeroesVI(t *testing.T) {
	target := 249
	result := trailingZeroes(1000)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
