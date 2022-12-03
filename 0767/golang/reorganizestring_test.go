package reorganizestring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReorganizeString(t *testing.T) {
	cases := []struct {
		wanted []string
		s      string
	}{
		{[]string{"aba"}, "aab"},
		{[]string{""}, "aaab"},
		{[]string{"aba"}, "aba"},
		{[]string{"abab", "baba"}, "aabb"},
		{[]string{"abaca", "acaba"}, "aaabc"},
		{[]string{"abcabc", "abacbc", "acabcb", "cbabac", "acbcba"}, "aabbcc"},
	}

	for _, c := range cases {
		t.Run(c.s, func(t *testing.T) {
			assert.Contains(t, c.wanted, reorganizeString(c.s))
		})
	}
}
