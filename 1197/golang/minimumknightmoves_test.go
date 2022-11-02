package minimumknightmoves

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMnKnightMoves(t *testing.T) {
	cases := []struct {
		wanted int
		x      int
		y      int
	}{
		{1, 2, 1},
		{4, 5, 5},
		{0, 0, 0},
		{68, 100, 100},
		{2, 1, 1},
		{94, -172, -110},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%d_%d", c.x, c.y), func(t *testing.T) {
			assert.Equal(t, c.wanted, minKnightMoves(c.x, c.y))
		})
	}
}
