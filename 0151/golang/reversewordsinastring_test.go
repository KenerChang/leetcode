package reversewordsinastring

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	input := "the sky is blue"
	target := "blue is sky the"

	result := reverseWords(input)
	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestReverseWordsII(t *testing.T) {
	input := "  hello world!  "
	target := "world! hello"

	result := reverseWords(input)
	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestReverseWordsIII(t *testing.T) {
	input := "a good   example"
	target := "example good a"

	result := reverseWords(input)
	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestReverseWordsIV(t *testing.T) {
	input := ""
	target := ""

	result := reverseWords(input)
	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestReverseWordsV(t *testing.T) {
	input := "   "
	target := ""

	result := reverseWords(input)
	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}
