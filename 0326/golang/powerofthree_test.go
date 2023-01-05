package powerofthree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfThree(t *testing.T) {
	cases := []struct {
		wanted bool
		n      int
	}{
		{true, 81},
		{true, 27},
		{false, 0},
		{true, 1},
		{false, -1},
		{false, 45},
		{false, 6},
		{false, 19684},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.n), func(t *testing.T) {
			assert.Equal(t, c.wanted, isPowerOfThree(c.n))
		})
	}
}
