package happynumber

import (
	"testing"
)

func TestIsHappy(t *testing.T) {
	target := true
	result := isHappy(19)

	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsHappyII(t *testing.T) {
	target := false
	result := isHappy(0)

	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsHappyIII(t *testing.T) {
	target := true
	result := isHappy(1)

	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}
