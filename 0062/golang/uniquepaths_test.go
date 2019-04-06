package uniquepaths

import (
	"testing"
)

func TestUniquePaths(t *testing.T) {
	m := 3
	n := 2
	target := 3

	result := uniquePaths(m, n)
	if target != result {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	m = 7
	n = 3
	target = 28

	result = uniquePaths(m, n)
	if target != result {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	m = 1
	n = 0
	target = 0

	result = uniquePaths(m, n)
	if target != result {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	m = 2
	n = 1
	target = 1

	result = uniquePaths(m, n)
	if target != result {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	m = 1
	n = 100
	target = 1

	result = uniquePaths(m, n)
	if target != result {
		t.Errorf("expect %d, but got: %d", target, result)
	}

	m = 2
	n = 100
	target = 100

	result = uniquePaths(m, n)
	if target != result {
		t.Errorf("expect %d, but got: %d", target, result)
	}

}
