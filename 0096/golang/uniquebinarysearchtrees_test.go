package uniquebinarysearchtrees

import (
	"testing"
)

func TestNumTrees(t *testing.T) {
	target := 5
	result := numTrees(3)

	if target != result {
		t.Errorf("expect %d, got: %d", target, result)
	}

	target = 1
	result = numTrees(0)
	if target != result {
		t.Errorf("expect %d, got: %d", target, result)
	}

	target = 1
	result = numTrees(1)
	if target != result {
		t.Errorf("expect %d, got: %d", target, result)
	}

	target = 2
	result = numTrees(2)
	if target != result {
		t.Errorf("expect %d, got: %d", target, result)
	}

	target = 16796
	result = numTrees(10)
	if target != result {
		t.Errorf("expect %d, got: %d", target, result)
	}

}
