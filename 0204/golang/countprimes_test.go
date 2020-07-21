package countprimes

import (
	"testing"
)

func TestCountPrimes(t *testing.T) {
	target := 4
	result := countPrimes(10)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCountPrimesII(t *testing.T) {
	target := 0
	result := countPrimes(2)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCountPrimesIII(t *testing.T) {
	target := 2
	result := countPrimes(5)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCountPrimesIV(t *testing.T) {
	target := 3
	result := countPrimes(7)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCountPrimesV(t *testing.T) {
	target := 41537
	result := countPrimes(499979)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
