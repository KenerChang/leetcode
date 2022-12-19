package firstuniquecharacterinastring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstUniqChar(t *testing.T) {
	cases := []struct {
		wanted int
		s      string
	}{
		{0, "leetcode"},
		{2, "loveleetcode"},
		{-1, "aabb"},
		{0, "l"},
		{0, "abc"},
		{2, "aabc"},
		{4, "aabbc"},
	}

	for _, c := range cases {
		t.Run(c.s, func(t *testing.T) {
			assert.Equal(t, c.wanted, firstUniqChar(c.s))
		})
	}
}
