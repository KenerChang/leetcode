package numberof1bits

import (
	"testing"
	"math"
)

func TestHammingWeight(t *testing.T) {
	var input uint32 = 11
	target := 3

	result := hammingWeight(input)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestHammingWeightII(t *testing.T) {
	var input uint32 = 128
	target := 1

	result := hammingWeight(input)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestHammingWeightIII(t *testing.T) {
	var input uint32 = 4294967293
	target := 31

	result := hammingWeight(input)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestHammingWeightIV(t *testing.T) {
	var input uint32 = 0
	target := 0

	result := hammingWeight(input)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestHammingWeightV(t *testing.T) {
	var input uint32 = math.MaxUint32
	target := 32

	result := hammingWeight(input)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
