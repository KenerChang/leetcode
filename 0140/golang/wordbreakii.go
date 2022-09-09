package wordbreakii

import (
	"strings"
)

type Trie struct {
	Children map[rune]*Trie
	IsEnd    bool
}

func (t *Trie) Insert(word string) {
	node := t
	for _, c := range word {
		if node.Children[c] == nil {
			node.Children[c] = &Trie{Children: make(map[rune]*Trie)}
		}
		node = node.Children[c]
	}
	node.IsEnd = true
}

func (t *Trie) Search(key string) bool {
	node := t
	for _, c := range key {
		if node.Children[c] == nil {
			return false
		}
		node = node.Children[c]
	}
	return node.IsEnd
}

func wordBreak(s string, wordDict []string) []string {
	// build trie
	root := &Trie{Children: make(map[rune]*Trie)}
	for _, word := range wordDict {
		root.Insert(word)
	}

	// solve word break by recursion with trie
	candidates := wordBreadImpl(s, root)

	results := []string{}
	for _, candidate := range candidates {
		results = append(results, strings.Join(candidate, " "))
	}
	return results
}

func wordBreadImpl(s string, trie *Trie) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}

	var candidates [][]string
	for i := 0; i < len(s); i++ {
		if trie.Search(s[:i+1]) {
			// s[0:i] is in dict

			if i == len(s)-1 {
				// reach end
				candidate := []string{s}
				candidates = append(candidates, candidate)
			} else {
				subCandiates := wordBreadImpl(s[i+1:], trie)
				for _, subCandiate := range subCandiates {
					candidate := append([]string{s[:i+1]}, subCandiate...)
					candidates = append(candidates, candidate)
				}
			}
		}
	}

	return candidates
}
