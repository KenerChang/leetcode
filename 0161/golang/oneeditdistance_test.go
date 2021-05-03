package oneeditdistance

import "testing"

func TestIsOneEditDistance(t *testing.T) {
	result := isOneEditDistance("ab", "acb")

	target := true
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsOneEditDistanceII(t *testing.T) {
	result := isOneEditDistance("", "")

	target := false
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsOneEditDistanceIII(t *testing.T) {
	result := isOneEditDistance("a", "")

	target := true
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsOneEditDistanceIV(t *testing.T) {
	result := isOneEditDistance("", "A")

	target := true
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsOneEditDistanceV(t *testing.T) {
	result := isOneEditDistance(
		"abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdef",
		"bcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefg",
	)

	target := false
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}
