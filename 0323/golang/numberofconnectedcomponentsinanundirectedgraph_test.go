package numberofconnectedcomponentsinanundirectedgraph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountComponents(t *testing.T) {
	cases := []struct {
		wanted int
		n      int
		edges  [][]int
	}{
		{2, 5, [][]int{{0, 1}, {1, 2}, {3, 4}}},
		{1, 5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}},
		{5, 5, [][]int{}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%d_%v", c.n, c.edges), func(t *testing.T) {
			assert.Equal(t, c.wanted, countComponents(c.n, c.edges))
		})
	}
}
