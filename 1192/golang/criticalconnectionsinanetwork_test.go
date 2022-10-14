package criticalconnectionsinanetwork

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCriticalConnections(t *testing.T) {
	cases := []struct {
		name        string
		wanted      [][]int
		n           int
		connections [][]int
	}{
		{"case 1", [][]int{{1, 3}}, 4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}},
		{"case 2", [][]int{{0, 1}}, 2, [][]int{{0, 1}}},
		{"case 3", [][]int{}, 3, [][]int{{0, 1}, {1, 2}, {2, 0}}},
		{"case 4", [][]int{{0, 1}, {1, 2}}, 3, [][]int{{0, 1}, {1, 2}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.wanted, criticalConnections(c.n, c.connections))
		})
	}
}
