package minimumheighttrees

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinHeightTrees(t *testing.T) {
	cases := []struct {
		wanted []int
		n      int
		edges  [][]int
	}{
		{[]int{1}, 4, [][]int{{1, 0}, {1, 2}, {1, 3}}},
		{[]int{3, 4}, 6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}},
		{[]int{1}, 3, [][]int{{0, 1}, {1, 2}}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%d_%v", c.n, c.edges), func(t *testing.T) {
			assert.ElementsMatch(t, c.wanted, findMinHeightTrees(c.n, c.edges))
		})
	}
}
