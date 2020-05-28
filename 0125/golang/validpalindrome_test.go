package validpalindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	target := true

	result := isPalindrome(s)
	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsPalindromeII(t *testing.T) {
	s := ""
	target := true

	result := isPalindrome(s)
	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsPalindromeIII(t *testing.T) {
	s := "aab      b    aa"
	target := true

	result := isPalindrome(s)
	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsPalindromeIV(t *testing.T) {
	s := "race a car"
	target := false

	result := isPalindrome(s)
	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsPalindromeV(t *testing.T) {
	s := "0P"
	target := false

	result := isPalindrome(s)
	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsPalindromeVI(t *testing.T) {
	s := "0P0"
	target := true

	result := isPalindrome(s)
	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}
