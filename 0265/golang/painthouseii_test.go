package painthouseii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostII(t *testing.T) {
	cases := []struct {
		costs  [][]int
		wanted int
	}{
		{[][]int{{1, 5, 3}, {2, 9, 4}}, 5},
		{[][]int{{1, 3}, {2, 4}}, 5},
		{[][]int{{1, 5, 4}, {2, 9, 4}, {2, 9, 4}}, 7},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("costs: %v", c.costs), func(t *testing.T) {
			assert.Equal(t, c.wanted, minCostII(c.costs))
		})
	}
}
