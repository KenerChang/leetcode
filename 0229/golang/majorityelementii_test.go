package majorityelementii

import "testing"

func verify(target, result []int) bool {
	if len(target) != len(result) {
		return false
	}

	for i, numt := range target {
		if numt != result[i] {
			return false
		}
	}
	return true
}

func TestMajorityElement(t *testing.T) {
	result := majorityElement([]int{3, 2, 3})

	target := []int{3}

	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestMajorityElementII(t *testing.T) {
	result := majorityElement([]int{1})

	target := []int{1}

	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestMajorityElementIII(t *testing.T) {
	result := majorityElement([]int{1, 2})

	target := []int{1, 2}

	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestMajorityElementIV(t *testing.T) {
	result := majorityElement([]int{3, 2, 3, 2, 2})

	target := []int{3, 2}

	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}
