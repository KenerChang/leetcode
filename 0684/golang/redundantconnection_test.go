package redundantconnection

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRedundantConnection(t *testing.T) {
	cases := []struct {
		wanted []int
		edges  [][]int
	}{
		{[]int{2, 3}, [][]int{{1, 2}, {1, 3}, {2, 3}}},
		{[]int{1, 4}, [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.edges), func(t *testing.T) {
			assert.Equal(t, c.wanted, findRedundantConnection(c.edges))
		})
	}
}
