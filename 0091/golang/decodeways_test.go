package decodeways

import (
	"testing"
)

func TestNumDecodings(t *testing.T) {
	template := "expect input: %s to be %d, got: %d"
	input := "12"
	target := 2

	result := numDecodings(input)
	if result != target {
		t.Errorf(template, input, target, result)
	}

	input = "02"
	target = 0

	result = numDecodings(input)
	if result != target {
		t.Errorf(template, input, target, result)
	}

	input = "0"
	target = 0

	result = numDecodings(input)
	if result != target {
		t.Errorf(template, input, target, result)
	}

	input = "22"
	target = 2

	result = numDecodings(input)
	if result != target {
		t.Errorf(template, input, target, result)
	}

	input = "222"
	target = 3

	result = numDecodings(input)
	if result != target {
		t.Errorf(template, input, target, result)
	}

	input = "22266"
	target = 5

	result = numDecodings(input)
	if result != target {
		t.Errorf(template, input, target, result)
	}

	input = "88"
	target = 1

	result = numDecodings(input)
	if result != target {
		t.Errorf(template, input, target, result)
	}

	input = "80"
	target = 0

	result = numDecodings(input)
	if result != target {
		t.Errorf(template, input, target, result)
	}

	input = "808"
	target = 0

	result = numDecodings(input)
	if result != target {
		t.Errorf(template, input, target, result)
	}

	input = "880"
	target = 0

	result = numDecodings(input)
	if result != target {
		t.Errorf(template, input, target, result)
	}

	input = "088"
	target = 0

	result = numDecodings(input)
	if result != target {
		t.Errorf(template, input, target, result)
	}

	input = "808"
	target = 0

	result = numDecodings(input)
	if result != target {
		t.Errorf(template, input, target, result)
	}

	input = "880"
	target = 0

	result = numDecodings(input)
	if result != target {
		t.Errorf(template, input, target, result)
	}
}
