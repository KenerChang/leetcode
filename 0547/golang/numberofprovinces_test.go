package numberofprovinces

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCircleNum(t *testing.T) {
	cases := []struct {
		wanted      int
		isConnected [][]int
	}{
		{2, [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}},
		{3, [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}},
		{1, [][]int{{0, 1, 0}, {1, 0, 1}, {0, 1, 0}}},
		{1, [][]int{{0}}},
		{1, [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.isConnected), func(t *testing.T) {
			assert.Equal(t, c.wanted, findCircleNum(c.isConnected))
		})
	}
}
