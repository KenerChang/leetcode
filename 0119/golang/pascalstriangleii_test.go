package pascalstriangleii

import (
	"testing"
)

func verify(target, result []int) bool {
	if len(target) != len(result) {
		return false
	}

	for i, num1 := range target {
		if num1 != result[i] {
			return false
		}
	}
	return true
}

func TestGetRow(t *testing.T) {
	target := []int{1, 3, 3, 1}
	result := getRow(3)

	if !verify(target, result) {
		t.Errorf("exepct %+v, got %+v", target, result)
	}
}

func TestGetRowII(t *testing.T) {
	target := []int{1}
	result := getRow(0)

	if !verify(target, result) {
		t.Errorf("exepct %+v, got %+v", target, result)
	}
}

func TestGetRowIII(t *testing.T) {
	target := []int{1, 1}
	result := getRow(1)

	if !verify(target, result) {
		t.Errorf("exepct %+v, got %+v", target, result)
	}
}

func TestGetRowIV(t *testing.T) {
	target := []int{1, 2, 1}
	result := getRow(2)

	if !verify(target, result) {
		t.Errorf("exepct %+v, got %+v", target, result)
	}
}

func TestGetRowV(t *testing.T) {
	target := []int{1, 4, 6, 4, 1}
	result := getRow(4)

	if !verify(target, result) {
		t.Errorf("exepct %+v, got %+v", target, result)
	}
}

func TestGetRowVI(t *testing.T) {
	target := []int{1, 5, 10, 10, 5, 1}
	result := getRow(5)

	if !verify(target, result) {
		t.Errorf("exepct %+v, got %+v", target, result)
	}
}
