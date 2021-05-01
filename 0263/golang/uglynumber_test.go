package uglynumber

import "testing"

func TestIsUgly(t *testing.T) {
	result := isUgly(6)

	target := true
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsUglyII(t *testing.T) {
	result := isUgly(14)

	target := false
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsUglyIII(t *testing.T) {
	result := isUgly(1)

	target := true
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsUglyIV(t *testing.T) {
	result := isUgly(20)

	target := true
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsUglyV(t *testing.T) {
	result := isUgly(0)

	target := false
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}
