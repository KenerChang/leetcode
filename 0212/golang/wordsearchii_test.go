package wordsearchii

import (
	"testing"
)

func verify(target, result []string) bool {
	if len(target) != len(result) {
		return false
	}

	targetMap, resultMap := map[string]bool{}, map[string]bool{}
	for _, s := range target {
		targetMap[s] = true
	}

	for _, s := range result {
		if _, found := targetMap[s]; !found {
			return false
		}
		resultMap[s] = true
	}

	for _, s := range target {
		if _, found := resultMap[s]; !found {
			return false
		}
	}
	return true
}

func TestFindWords(t *testing.T) {
	board := [][]byte{
		[]byte{'o', 'a', 'a', 'n'},
		[]byte{'e', 't', 'a', 'e'},
		[]byte{'i', 'h', 'k', 'r'},
		[]byte{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}

	target := []string{"eat", "oath"}
	result := findWords(board, words)
	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestFindWordsII(t *testing.T) {
	board := [][]byte{
		[]byte{'a', 'a'},
	}
	words := []string{"aaa"}

	target := []string{}
	result := findWords(board, words)
	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestFindWordsIII(t *testing.T) {
	// ["a","b","c"],["a","e","d"],["a","f","g"]
	board := [][]byte{
		[]byte{'a', 'b', 'c'},
		[]byte{'a', 'e', 'd'},
		[]byte{'a', 'f', 'g'},
	}
	words := []string{"abcdefg", "gfedcbaaa", "eaabcdgfa", "befa", "dgc", "ade"}

	target := []string{"abcdefg", "befa", "eaabcdgfa", "gfedcbaaa"}
	result := findWords(board, words)
	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestFindWordsIV(t *testing.T) {
	// ["a","b","c"],["a","e","d"],["a","f","g"]
	board := [][]byte{
		[]byte{'a'},
	}
	words := []string{"a"}

	target := []string{"a"}
	result := findWords(board, words)
	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}
