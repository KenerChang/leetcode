package findcriticalandpseudociticaledgesinminimumspanningtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCriticalAndPseudoCriticalEdges(t *testing.T) {
	cases := []struct {
		name   string
		wanted [][]int
		n      int
		edges  [][]int
	}{
		{"case 1", [][]int{{0, 1}, {2, 3, 4, 5}}, 5, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {0, 4, 3}, {3, 4, 3}, {1, 4, 6}}},
		{"case 2", [][]int{{}, {0, 1, 2, 3}}, 4, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 1}}},
		{"case 3", [][]int{{}, {0, 1, 2, 3, 4, 5}}, 4, [][]int{{0, 1, 1}, {0, 3, 1}, {0, 2, 1}, {1, 2, 1}, {1, 3, 1}, {2, 3, 1}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, findCriticalAndPseudoCriticalEdges(c.n, c.edges))
		})
	}
}
