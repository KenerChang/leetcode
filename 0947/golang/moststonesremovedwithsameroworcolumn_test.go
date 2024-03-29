package moststonesremovedwithsameroworcolumn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveStones(t *testing.T) {
	cases := []struct {
		wanted int
		stones [][]int
	}{
		{
			5,
			[][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}},
		},
		{
			3,
			[][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}},
		},
		{
			0,
			[][]int{{0, 0}},
		},
		{
			0,
			[][]int{{0, 0}, {1, 1}},
		},
		{
			1,
			[][]int{{0, 0}, {0, 1}},
		},
		{
			1,
			[][]int{{0, 0}, {1, 0}},
		},
		{
			2,
			[][]int{{0, 1}, {1, 0}, {1, 1}},
		},
		{
			5,
			[][]int{{0, 1}, {1, 2}, {1, 3}, {3, 3}, {2, 3}, {0, 2}},
		},
		{
			6,
			[][]int{{3, 3}, {4, 4}, {1, 4}, {1, 5}, {2, 3}, {4, 3}, {2, 4}},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.stones), func(t *testing.T) {
			assert.Equal(t, c.wanted, removeStones(c.stones))
		})
	}
}
