package furthestbuildingyoucanreach

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFurthestBuilding(t *testing.T) {
	cases := []struct {
		wanted  int
		heights []int
		bricks  int
		ladders int
	}{
		{4, []int{4, 2, 7, 6, 9, 14, 12}, 5, 1},
		{7, []int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2},
		{3, []int{14, 3, 19, 3}, 17, 0},
		{3, []int{5, 4, 3, 2}, 0, 0},
		{0, []int{2, 3, 4, 5}, 0, 0},
		{1, []int{2, 3, 4, 5}, 0, 1},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v@@%d@@%d", c.heights, c.bricks, c.ladders), func(t *testing.T) {
			assert.Equal(t, c.wanted, furthestBuilding(c.heights, c.bricks, c.ladders))
		})
	}
}
