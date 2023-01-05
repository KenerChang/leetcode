package eliminationgame

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastRemaining(t *testing.T) {
	cases := []struct {
		wanted int
		n      int
	}{
		{6, 9},
		{1, 1},
		{2, 2},
		{2, 4},
		{8, 10},
		{534765398, 1000000000},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.n), func(t *testing.T) {
			assert.Equal(t, c.wanted, lastRemaining(c.n))
		})
	}
}
