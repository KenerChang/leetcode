package majorityelement

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	target := 3
	result := majorityElement([]int{3, 2, 3})

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMajorityElementII(t *testing.T) {
	target := 2
	result := majorityElement([]int{2, 2, 1, 1, 1, 2, 2})

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
