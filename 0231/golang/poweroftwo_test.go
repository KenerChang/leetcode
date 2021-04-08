package poweroftwo

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	result := isPowerOfTwo(1)

	target := true
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsPowerOfTwoII(t *testing.T) {
	result := isPowerOfTwo(16)

	target := true
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsPowerOfTwoIII(t *testing.T) {
	result := isPowerOfTwo(0)

	target := false
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsPowerOfTwoIV(t *testing.T) {
	result := isPowerOfTwo(15)

	target := false
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}
