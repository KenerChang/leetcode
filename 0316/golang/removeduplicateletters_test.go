package removeduplicateletters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicateLetters(t *testing.T) {
	cases := []struct {
		wanted string
		s      string
	}{
		{"abc", "bcabc"},
		{"acdb", "cbacdcbc"},
		{"bca", "bca"},
		{"dabc", "dcabc"},
		{"a", "a"},
		{"abc", "aabbcc"},
		{"abc", "cbaabc"},
		{"az", "zaz"},
		{"bac", "bbcaac"},
		{"bac", "bbbcaccc"},
		{"bac", "bbbcccac"},
	}

	for _, c := range cases {
		t.Run(c.s, func(t *testing.T) {
			assert.Equal(t, c.wanted, removeDuplicateLetters(c.s))
		})
	}
}
