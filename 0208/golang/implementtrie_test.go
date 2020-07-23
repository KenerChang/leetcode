package implementtrie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := Constructor()

	trie.Insert("apple")

	if !trie.Search("apple") {
		t.Errorf("expect true, got false")
	}

	if trie.Search("app") {
		t.Errorf("expect false, got true")
	}

	if !trie.StartsWith("app") {
		t.Errorf("expect true, got false")
	}

	trie.Insert("app")

	if !trie.Search("app") {
		t.Errorf("expect true, got false")
	}
}
