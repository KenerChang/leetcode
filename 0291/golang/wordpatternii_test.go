package wordpatternii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordPatternMatch(t *testing.T) {
	cases := []struct {
		wanted  bool
		pattern string
		s       string
	}{
		{true, "abab", "redblueredblue"},
		{true, "aaaa", "asdasdasdasd"},
		{false, "aabb", "xyzabcxzyabc"},
		{true, "a", "red"},
		{true, "aba", "redr"},
		{false, "aba", "redc"},
		{false, "gh", "i"},
		{false, "ab", "aa"},
		{true, "aa", "aa"},
		{false, "aa", "ac"},
		{true, "sucks", "teezmnoteez"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("pattern:%s,s:%s", c.pattern, c.s), func(t *testing.T) {
			assert.Equal(t, c.wanted, wordPatternMatch(c.pattern, c.s))
		})
	}
}
