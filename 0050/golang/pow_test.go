package pow

import (
	"testing"
)

func TestMyPow(t *testing.T) {
	x := 2.00
	n := 10
	target := 1024.00
	result := myPow(x, n)
	if target != result {
		t.Errorf("expect %f but got %f", target, result)
	}

	x = 4.00
	n = 10
	target = 1048576
	result = myPow(x, n)
	if target != result {
		t.Errorf("expect %f but got %f", target, result)
	}

	x = 1.00
	n = 10
	target = 1
	result = myPow(x, n)
	if target != result {
		t.Errorf("expect %f but got %f", target, result)
	}

	x = 0.00
	n = 147483647
	target = 0
	result = myPow(x, n)
	if target != result {
		t.Errorf("expect %f but got %f", target, result)
	}

	x = 2.00
	n = 5
	target = 32
	result = myPow(x, n)
	if target != result {
		t.Errorf("expect %f but got %f", target, result)
	}

	x = 100.00
	n = 2
	target = 10000.00
	result = myPow(x, n)
	if target != result {
		t.Errorf("expect %f but got %f", target, result)
	}

	x = 2.00
	n = -2
	target = 0.25
	result = myPow(x, n)
	if target != result {
		t.Errorf("expect %f but got %f", target, result)
	}

	x = 2.00
	n = 2
	target = 4.00
	result = myPow(x, n)
	if target != result {
		t.Errorf("expect %f but got %f", target, result)
	}

	x = 8.95371
	n = -1
	target = float64(1) / 8.95371
	result = myPow(x, n)
	if target != result {
		t.Errorf("expect %f but got %f", target, result)
	}
}
