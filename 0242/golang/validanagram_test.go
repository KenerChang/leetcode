package validanagram

import "testing"

func TestIsAnagram(t *testing.T) {
	result := isAnagram("anagram", "nagaram")

	target := true
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsAnagramII(t *testing.T) {
	result := isAnagram("rat", "car")

	target := false
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsAnagramIII(t *testing.T) {
	result := isAnagram("r", "r")

	target := true
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}
