package wordbreakii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordBreak(t *testing.T) {
	cases := []struct {
		name     string
		wanted   []string
		s        string
		wordDict []string
	}{
		{"case 1", []string{"cats and dog", "cat sand dog"}, "catsanddog", []string{"cat", "cats", "and", "sand", "dog"}},
		{"case 2", []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"}, "pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}},
		{"case 3", []string{}, "catsandog", []string{"cats", "dog", "sand", "and", "cat"}},
		{"case 4", []string{"a a a a a a a", "aa a a a a a", "a aa a a a a", "a a aa a a a", "aa aa a a a", "aaaa a a a", "a a a aa a a", "aa a aa a a", "a aa aa a a", "a aaaa a a", "a a a a aa a", "aa a a aa a", "a aa a aa a", "a a aa aa a", "aa aa aa a", "aaaa aa a", "a a aaaa a", "aa aaaa a", "a a a a a aa", "aa a a a aa", "a aa a a aa", "a a aa a aa", "aa aa a aa", "aaaa a aa", "a a a aa aa", "aa a aa aa", "a aa aa aa", "a aaaa aa", "a a a aaaa", "aa a aaaa", "a aa aaaa"}, "aaaaaaa", []string{"aaaa", "aa", "a"}},
	}

	for _, c := range cases {
		assert.ElementsMatch(t, c.wanted, wordBreak(c.s, c.wordDict))
	}
}
