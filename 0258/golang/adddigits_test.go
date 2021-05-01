package adddigits

import "testing"

func TestAddDigits(t *testing.T) {
	result := addDigits(38)

	target := 2
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestAddDigitsII(t *testing.T) {
	result := addDigits(0)

	target := 0
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestAddDigitsIII(t *testing.T) {
	result := addDigits(111)

	target := 3
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestAddDigitsIV(t *testing.T) {
	result := addDigits(5)

	target := 5
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}
