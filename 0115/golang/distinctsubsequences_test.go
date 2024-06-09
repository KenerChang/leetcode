package distinctsubsequences

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumDistinct(t *testing.T) {
	cases := []struct {
		s      string
		t      string
		wanted int
	}{
		{
			s:      "rabbbit",
			t:      "rabbit",
			wanted: 3,
		},
		{
			s:      "babgbag",
			t:      "bag",
			wanted: 5,
		},
		{
			s:      "babgbag",
			t:      "b",
			wanted: 3,
		},
		{
			s:      "rabbit",
			t:      "rabbbit",
			wanted: 0,
		},
		{
			s:      "r",
			t:      "ra",
			wanted: 0,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("s: %s, t: %s", c.s, c.t), func(t *testing.T) {
			assert.Equal(t, c.wanted, numDistinct(c.s, c.t))
		})
	}
}
