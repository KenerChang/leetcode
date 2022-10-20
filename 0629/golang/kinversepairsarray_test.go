package kinversepairsarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKInversePairs(t *testing.T) {
	cases := []struct {
		wanted int
		n      int
		k      int
	}{
		{1, 3, 0},
		{2, 3, 1},
		{955735232, 1000, 500},
		{0, 3, 5},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%d_%d", c.n, c.k), func(t *testing.T) {
			assert.Equal(t, c.wanted, kInversePairs(c.n, c.k))
		})
	}
}
