package mincosttoconnectallpoints

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostConnectPoints(t *testing.T) {
	cases := []struct {
		name   string
		wanted int
		given  [][]int
	}{
		{"case 1", 20, [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}},
		{"case 2", 18, [][]int{{3, 12}, {-2, 5}, {-4, 1}}},
		{"case 3", 4, [][]int{{0, 0}, {1, 1}, {1, 0}, {-1, 1}}},
		{"case 4", 4000000, [][]int{{-1000000, -1000000}, {1000000, 1000000}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, minCostConnectPoints(c.given))
		})
	}
}
