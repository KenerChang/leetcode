package removeduplicatedsortedarray

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	template := "expect %d, got %d"
	input := []int{1, 1, 1, 2, 2, 3}
	target := 5

	result := removeDuplicates(input)
	if target != result {
		t.Errorf(template, target, result)
	}

	input = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	target = 7

	result = removeDuplicates(input)
	if target != result {
		t.Errorf(template, target, result)
	}

	input = []int{1, 1, 1, 2, 2, 2, 3, 3}
	target = 6

	result = removeDuplicates(input)
	if target != result {
		t.Errorf(template, target, result)
	}

}
