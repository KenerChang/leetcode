package longestduplicatesubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestDupSubstring(t *testing.T) {
	cases := []struct {
		name   string
		wanted string
		s      string
	}{
		{"banana", "ana", "banana"},
		{"abcd", "", "abcd"},
		{"abcda", "a", "abcda"},
		{"aaaaa", "aaaa", "aaaaa"},
		{"aaabaaa", "aaa", "aaabaaa"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, longestDupSubstring(c.s))
		})
	}
}
