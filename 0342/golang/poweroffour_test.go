package poweroffour

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfFour(t *testing.T) {
	cases := []struct {
		wanted bool
		n      int
	}{
		{true, 16},
		{false, 5},
		{true, 1},
		{false, 0},
		{false, -1},
		{false, 6},
		{false, 12},
		{false, 8},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.n), func(t *testing.T) {
			assert.Equal(t, c.wanted, isPowerOfFour(c.n))
		})
	}
}
