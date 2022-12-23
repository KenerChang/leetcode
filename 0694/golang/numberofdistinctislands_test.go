package numberofdistinctislands

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumDistinctIslands(t *testing.T) {
	cases := []struct {
		wanted int
		grids  [][]int
	}{
		{1, [][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}},
		{3, [][]int{{1, 1, 0, 1, 1}, {1, 0, 0, 0, 0}, {0, 0, 0, 0, 1}, {1, 1, 0, 1, 1}}},
		{2, [][]int{{1, 1, 0, 1, 1}, {1, 0, 0, 0, 0}, {0, 0, 0, 1, 1}, {1, 1, 0, 1, 0}}},
		{1, [][]int{{1}}},
		{0, [][]int{{0}}},
		{1, [][]int{{1, 0, 0, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}, {1, 0, 0, 1}}},
		{2, [][]int{{1, 1, 0}, {0, 1, 1}, {0, 0, 0}, {1, 1, 1}, {0, 1, 0}}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.grids), func(t *testing.T) {
			assert.Equal(t, c.wanted, numDistinctIslands(c.grids))
		})
	}
}
