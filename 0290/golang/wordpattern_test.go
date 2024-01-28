package wordpattern

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordPattern(t *testing.T) {
	cases := []struct {
		pattern string
		s       string
		wanted  bool
	}{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("pattern: %s, s: %s", c.pattern, c.s), func(t *testing.T) {
			assert.Equal(t, c.wanted, wordPattern(c.pattern, c.s))
		})
	}
}
