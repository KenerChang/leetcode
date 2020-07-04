package largestnumber

import (
	"testing"
)

func TestLargestNumber(t *testing.T) {
	input := []int{10, 2}
	target := "210"

	result := largestNumber(input)
	if target != result {
		t.Errorf("exepct %s, got %s", target, result)
	}
}

func TestLargestNumberII(t *testing.T) {
	input := []int{3, 30, 34, 5, 9}
	target := "9534330"

	result := largestNumber(input)
	if target != result {
		t.Errorf("exepct %s, got %s", target, result)
	}
}

func TestLargestNumberIII(t *testing.T) {
	input := []int{1}
	target := "1"

	result := largestNumber(input)
	if target != result {
		t.Errorf("exepct %s, got %s", target, result)
	}
}

func TestLargestNumberIV(t *testing.T) {
	input := []int{1, 0}
	target := "10"

	result := largestNumber(input)
	if target != result {
		t.Errorf("exepct %s, got %s", target, result)
	}
}

func TestLargestNumberV(t *testing.T) {
	input := []int{}
	target := ""

	result := largestNumber(input)
	if target != result {
		t.Errorf("exepct %s, got %s", target, result)
	}
}

func TestLargestNumberVI(t *testing.T) {
	input := []int{0, 0}
	target := "0"

	result := largestNumber(input)
	if target != result {
		t.Errorf("exepct %s, got %s", target, result)
	}
}

func TestLargestNumberVII(t *testing.T) {
	input := []int{1, 1}
	target := "11"

	result := largestNumber(input)
	if target != result {
		t.Errorf("exepct %s, got %s", target, result)
	}
}

func TestLargestNumberVIII(t *testing.T) {
	input := []int{0, 0, 0}
	target := "0"

	result := largestNumber(input)
	if target != result {
		t.Errorf("exepct %s, got %s", target, result)
	}
}

func TestLargestNumberIX(t *testing.T) {
	input := []int{1, 1, 0, 0, 0}
	target := "11000"

	result := largestNumber(input)
	if target != result {
		t.Errorf("exepct %s, got %s", target, result)
	}
}
