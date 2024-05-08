package scramblestring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsScramble(t *testing.T) {
	cases := []struct {
		s1     string
		s2     string
		wanted bool
	}{
		{
			s1:     "great",
			s2:     "rgeat",
			wanted: true,
		},
		{
			s1:     "great",
			s2:     "atrge",
			wanted: true,
		},
		{
			s1:     "great",
			s2:     "atgre",
			wanted: true,
		},
		{
			s1:     "great",
			s2:     "tagre",
			wanted: true,
		},
		{
			s1:     "gr",
			s2:     "rg",
			wanted: true,
		},
		{
			s1:     "aa",
			s2:     "aa",
			wanted: true,
		},
		{
			s1:     "ab",
			s2:     "ba",
			wanted: true,
		},
		{
			s1:     "abcde",
			s2:     "caebd",
			wanted: false,
		},
		{
			s1:     "a",
			s2:     "a",
			wanted: true,
		},
		{
			s1:     "a",
			s2:     "b",
			wanted: false,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("s1: %s, s2: %s", c.s1, c.s2), func(t *testing.T) {
			assert.Equal(t, c.wanted, isScramble(c.s1, c.s2))
		})
	}
}
