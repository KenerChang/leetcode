package reversebits

import (
	"math"
	"testing"
)

func TestReverseBits(t *testing.T) {
	var input uint32 = 43261596
	var target uint32 = 964176192

	result := reverseBits(input)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestReverseBitsII(t *testing.T) {
	var input uint32 = 4294967293
	var target uint32 = 3221225471

	result := reverseBits(input)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestReverseBitsIII(t *testing.T) {
	var input uint32 = math.MaxUint32
	var target uint32 = math.MaxUint32

	result := reverseBits(input)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
