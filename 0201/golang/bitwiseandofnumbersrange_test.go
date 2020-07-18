package bitwiseandofnumbersrange

import (
	"testing"
)

func TestRangeBitwiseAnd(t *testing.T) {
	target := 4
	result := rangeBitwiseAnd(5, 7)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestRangeBitwiseAndII(t *testing.T) {
	target := 0
	result := rangeBitwiseAnd(0, 1)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestRangeBitwiseAndIII(t *testing.T) {
	target := 0
	result := rangeBitwiseAnd(1, 2147483647)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestRangeBitwiseAndIV(t *testing.T) {
	target := 0
	result := rangeBitwiseAnd(33, 127)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestRangeBitwiseAndV(t *testing.T) {
	target := 32
	result := rangeBitwiseAnd(33, 63)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestRangeBitwiseAndVI(t *testing.T) {
	target := 6
	result := rangeBitwiseAnd(6, 7)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestRangeBitwiseAndVII(t *testing.T) {
	target := 60
	result := rangeBitwiseAnd(60, 63)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestRangeBitwiseAndVIII(t *testing.T) {
	target := 0
	result := rangeBitwiseAnd(3, 6)

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
