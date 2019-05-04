package searchinrotatedsortedarray

import (
	"testing"
)

func TestSearch(t *testing.T) {
	template := "expect %+v and %d to be %t, got: %t"
	input := []int{2, 5, 6, 0, 0, 1, 2}
	num := 0
	target := true

	result := search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	num = 2
	target = true

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	num = 8
	target = false

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	input = []int{3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 7, 0, 1, 2, 2}
	num = 3
	target = true

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	num = 2
	target = true

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	num = 6
	target = true

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	num = 8
	target = false

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	input = []int{1}
	num = 1
	target = true

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	num = 0
	target = false

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	input = []int{1, 2}
	num = 1
	target = true

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	num = 2
	target = true

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	num = 0
	target = false

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	input = []int{3, 1, 2}
	num = 1
	target = true

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	num = 2
	target = true

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	num = 3
	target = true

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	num = 0
	target = false

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	input = []int{1, 1, 3, 1}
	num = 3
	target = true

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	num = 1
	target = true

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	num = 2
	target = false

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	input = []int{7, 8, 9, 1, 2, 3, 4, 5, 6, 7}
	num = 8
	target = true

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	num = 9
	target = true

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	num = 7
	target = true

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	num = 3
	target = true

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	num = 10
	target = false

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}

	input = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	num = 2
	target = true

	result = search(input, num)
	if target != result {
		t.Errorf(template, input, num, target, result)
	}
}
