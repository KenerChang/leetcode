package wordladder

import (
	"testing"
)

func TestLadderLength(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	target := 5
	result := ladderLength(beginWord, endWord, wordList)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestLadderLengthII(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log"}

	target := 0
	result := ladderLength(beginWord, endWord, wordList)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestLadderLengthIII(t *testing.T) {
	beginWord := "hot"
	endWord := "dog"
	wordList := []string{"hot", "dog"}

	target := 0
	result := ladderLength(beginWord, endWord, wordList)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestLadderLengthIV(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"kit", "dot", "kip", "kop", "kog", "log", "cog"}

	target := 6
	result := ladderLength(beginWord, endWord, wordList)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestLadderLengthV(t *testing.T) {
	beginWord := "qa"
	endWord := "sq"
	wordList := []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}

	target := 5
	result := ladderLength(beginWord, endWord, wordList)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestLadderLengthVI(t *testing.T) {
	beginWord := "game"
	endWord := "thee"
	wordList := []string{"frye", "heat", "tree", "thee", "free", "hell", "fame", "faye"}

	target := 7
	result := ladderLength(beginWord, endWord, wordList)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
