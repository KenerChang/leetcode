package substringwithlargestvariance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestVariance(t *testing.T) {
	cases := []struct {
		wanted int
		s      string
	}{
		{3, "aababbb"},
		{0, "abcde"},
	}

	for _, c := range cases {
		t.Run(c.s, func(t *testing.T) {
			assert.Equal(t, c.wanted, largestVariance(c.s))
		})
	}
}
