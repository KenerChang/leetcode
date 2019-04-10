package sqrt

import (
	"testing"
)

func TestMySqrt(t *testing.T) {
	input := 4
	target := 2

	result := mySqrt(input)
	if result != target {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	input = 8
	target = 2

	result = mySqrt(input)
	if result != target {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	input = 1
	target = 1

	result = mySqrt(input)
	if result != target {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	input = 0
	target = 0

	result = mySqrt(input)
	if result != target {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	input = 100
	target = 10

	result = mySqrt(input)
	if result != target {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	input = 1000
	target = 31

	result = mySqrt(input)
	if result != target {
		t.Errorf("expect %d, but got: %d", target, result)
	}

}
