package numberofislandsii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumIslands2(t *testing.T) {
	cases := []struct {
		m         int
		n         int
		positions [][]int
		wanted    []int
	}{
		{3, 3, [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}}, []int{1, 1, 2, 3}},
		{1, 1, [][]int{{1, 1}}, []int{1}},
		{3, 3, [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}, {1, 1}}, []int{1, 1, 2, 3, 1}},
		{3, 3, [][]int{{0, 0}, {1, 1}, {0, 1}, {1, 2}, {2, 1}}, []int{1, 2, 1, 1, 1}},
		{3, 3, [][]int{{0, 1}, {1, 2}, {2, 1}, {1, 0}, {0, 2}, {0, 0}, {1, 1}}, []int{1, 2, 3, 4, 3, 2, 1}},
		{3, 3, [][]int{{0, 0}, {0, 1}, {1, 2}, {1, 2}}, []int{1, 1, 2, 2}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("m:%d, n:%d, positiions: %v", c.m, c.n, c.positions), func(t *testing.T) {
			assert.Equal(t, c.wanted, numIslands2(c.m, c.n, c.positions))
		})
	}
}
