package longestsubstringwithoutrepeatingcharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		target int
		s      string
	}{
		{
			3, "abcabcbb",
		},
		{
			1, "bbbbb",
		},
		{
			3, "pwwkew",
		},
		{
			7, "abc abcd  abcdef",
		},
	}

	for _, c := range cases {
		t.Run(c.s, func(t *testing.T) {
			assert.Equal(t, c.target, lengthOfLongestSubstring(c.s))
		})
	}
}
