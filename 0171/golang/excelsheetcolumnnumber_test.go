package excelsheetcolumnnumber

import (
	"testing"
)

func TestTitleToNumber(t *testing.T) {
	target := 1
	result := titleToNumber("A")

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestTitleToNumberII(t *testing.T) {
	target := 28
	result := titleToNumber("AB")

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestTitleToNumberIII(t *testing.T) {
	target := 701
	result := titleToNumber("ZY")

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestTitleToNumberIV(t *testing.T) {
	target := 52
	result := titleToNumber("AZ")

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestTitleToNumberV(t *testing.T) {
	target := 28
	result := titleToNumber("AB")

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestTitleToNumberVI(t *testing.T) {
	target := 26
	result := titleToNumber("Z")

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
