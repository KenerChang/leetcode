package compareversionnumbers

import (
	"testing"
)

func TestCompareVersion(t *testing.T) {
	target := -1
	result := compareVersion("0.1", "1.1")

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCompareVersionII(t *testing.T) {
	target := 1
	result := compareVersion("1.0.1", "1")

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCompareVersionIII(t *testing.T) {
	target := -1
	result := compareVersion("7.5.2.4", "7.5.3")

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCompareVersionIV(t *testing.T) {
	target := 0
	result := compareVersion("1.01", "1.001")

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCompareVersionV(t *testing.T) {
	target := 0
	result := compareVersion("1.0", "1.0.0")

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCompareVersionVI(t *testing.T) {
	target := 1
	result := compareVersion("1.15", "1.9.0")

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCompareVersionVII(t *testing.T) {
	target := -1
	result := compareVersion(".1", "1.1")

	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
