package minimumdeletionstomakecharacterfrequenciesunique

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDeletions(t *testing.T) {
	cases := []struct {
		wanted int
		s      string
	}{
		{0, "aab"},
		{2, "aaabbbcc"},
		{2, "ceabaacb"},
	}

	for _, c := range cases {
		t.Run(c.s, func(t *testing.T) {
			assert.Equal(t, c.wanted, minDeletions(c.s))
		})
	}
}
