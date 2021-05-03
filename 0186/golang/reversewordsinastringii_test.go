package reversewordsinastringii

import "testing"

func verify(target, result []byte) bool {
	if len(target) != len(result) {
		return false
	}

	for i := 0; i < len(target); i++ {
		if target[i] != result[i] {
			return false
		}
	}
	return true
}

func TestReverseWords(t *testing.T) {
	input := []byte{'t', 'h', 'e', ' ', 's', 'k', 'y', ' ', 'i', 's', ' ', 'b', 'l', 'u', 'e'}
	reverseWords(input)

	target := []byte{'b', 'l', 'u', 'e', ' ', 'i', 's', ' ', 's', 'k', 'y', ' ', 't', 'h', 'e'}
	if !verify(target, input) {
		t.Errorf("expect %v, got %v", target, input)
	}
}

func TestReverseWordsII(t *testing.T) {
	input := []byte{'a'}
	reverseWords(input)

	target := []byte{'a'}
	if !verify(target, input) {
		t.Errorf("expect %v, got %v", target, input)
	}
}
