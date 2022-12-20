package longestsubstringwithatmosttwodistinctcharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstringTwoDistinct(t *testing.T) {
	cases := []struct {
		wanted int
		s      string
	}{
		{3, "eceba"},
		{5, "ccaabbb"},
		{3, "eee"},
		{6, "aaAAAbbbBBB"},
		{2, "abcabc"},
	}

	for _, c := range cases {
		t.Run(c.s, func(t *testing.T) {
			assert.Equal(t, c.wanted, lengthOfLongestSubstringTwoDistinct(c.s))
		})
	}
}
