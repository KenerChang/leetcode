package isomorphicstrings

import (
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	target := true
	result := isIsomorphic("egg", "add")

	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsIsomorphicII(t *testing.T) {
	target := false
	result := isIsomorphic("foo", "bar")

	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsIsomorphicIII(t *testing.T) {
	target := true
	result := isIsomorphic("paper", "title")

	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsIsomorphicIV(t *testing.T) {
	target := true
	result := isIsomorphic("paper", "paper")

	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsIsomorphicV(t *testing.T) {
	target := false
	result := isIsomorphic("ab", "aa")

	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsIsomorphicVI(t *testing.T) {
	target := false
	result := isIsomorphic("aa", "ab")

	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}
