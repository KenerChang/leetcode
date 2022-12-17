package generalizedabbreviation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAbbreviations(t *testing.T) {
	cases := []struct {
		wanted []string
		word   string
	}{
		{[]string{"4", "3d", "2r1", "2rd", "1o2", "1o1d", "1or1", "1ord", "w3", "w2d", "w1r1", "w1rd", "wo2", "wo1d", "wor1", "word"}, "word"},
		{[]string{"1", "a"}, "a"},
		{[]string{"2", "1a", "a1", "aa"}, "aa"},
		{[]string{"6", "5f", "4e1", "4ef", "3d2", "3d1f", "3de1", "3def", "2c3", "2c2f", "2c1e1", "2c1ef", "2cd2", "2cd1f", "2cde1", "2cdef", "1b4", "1b3f", "1b2e1", "1b2ef", "1b1d2", "1b1d1f", "1b1de1", "1b1def", "1bc3", "1bc2f", "1bc1e1", "1bc1ef", "1bcd2", "1bcd1f", "1bcde1", "1bcdef", "a5", "a4f", "a3e1", "a3ef", "a2d2", "a2d1f", "a2de1", "a2def", "a1c3", "a1c2f", "a1c1e1", "a1c1ef", "a1cd2", "a1cd1f", "a1cde1", "a1cdef", "ab4", "ab3f", "ab2e1", "ab2ef", "ab1d2", "ab1d1f", "ab1de1", "ab1def", "abc3", "abc2f", "abc1e1", "abc1ef", "abcd2", "abcd1f", "abcde1", "abcdef"}, "abcdef"},
	}

	for _, c := range cases {
		t.Run(c.word, func(t *testing.T) {
			assert.ElementsMatch(t, c.wanted, generateAbbreviations(c.word))
		})
	}
}
