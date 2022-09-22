package repeatedstringmatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatedStringMatch(t *testing.T) {
	cases := []struct {
		name   string
		wanted int
		a, b   string
	}{
		{"case 1", 3, "abcd", "cdabcdab"},
		{"case 2", 2, "a", "aa"},
		{"case 3", 1, "a", "a"},
		{"case 4", 4, "abc", "cabcabca"},
		{"case 5", 2, "aaaaaaaaaaaaaaaaaaaaaab", "ba"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, repeatedStringMatch(c.a, c.b))
		})
	}
}
