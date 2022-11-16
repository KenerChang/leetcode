package rottingoranges

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrangesRotting(t *testing.T) {
	cases := []struct {
		wanted int
		grid   [][]int
	}{
		{4, [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}},
		{-1, [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}},
		{0, [][]int{{2}}},
		{-1, [][]int{{1}}},
		{0, [][]int{{0}}},
		{1, [][]int{{2, 1}}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.grid), func(t *testing.T) {
			assert.Equal(t, c.wanted, orangesRotting(c.grid))
		})
	}
}
