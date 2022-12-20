package longestsubstringwithatmostkdistinctcharacters

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstringKDistinct(t *testing.T) {
	cases := []struct {
		wanted int
		s      string
		k      int
	}{
		{3, "eceba", 2},
		{2, "aa", 1},
		{0, "eceba", 0},
		{1, "abcd", 1},
		{3, "aaa", 1},
		{1, "abacd", 1},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("s:%s,k:%d", c.s, c.k), func(t *testing.T) {
			assert.Equal(t, c.wanted, lengthOfLongestSubstringKDistinct(c.s, c.k))
		})
	}
}
