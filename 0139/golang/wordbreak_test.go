package wordbreak

import (
	"testing"
)

func TestWordBreak(t *testing.T) {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	target := true

	result := wordBreak(s, wordDict)
	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestWordBreakII(t *testing.T) {
	s := "applepenapple"
	wordDict := []string{"apple", "pen"}
	target := true

	result := wordBreak(s, wordDict)
	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestWordBreakIII(t *testing.T) {
	s := "catsandog"
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	target := false

	result := wordBreak(s, wordDict)
	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}
