package shortestpalindrome

import (
	"testing"
)

func TestShortestPalindrome(t *testing.T) {
	input := "aacecaaa"
	target := "aaacecaaa"

	result := shortestPalindrome(input)
	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestShortestPalindromeII(t *testing.T) {
	input := ""
	target := ""

	result := shortestPalindrome(input)
	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestShortestPalindromeIII(t *testing.T) {
	input := "a"
	target := "a"

	result := shortestPalindrome(input)
	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestShortestPalindromeIV(t *testing.T) {
	input := "abcdefg"
	target := "gfedcbabcdefg"

	result := shortestPalindrome(input)
	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestShortestPalindromeV(t *testing.T) {
	input := "ab"
	target := "bab"

	result := shortestPalindrome(input)
	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestShortestPalindromeVI(t *testing.T) {
	input := "bacbb"
	target := "bbcabacbb"

	result := shortestPalindrome(input)
	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestShortestPalindromeVII(t *testing.T) {
	input := "caab"
	target := "baacaab"

	result := shortestPalindrome(input)
	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}
