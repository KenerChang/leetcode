package longestincreasingpathinamatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestIncreasingPath(t *testing.T) {
	cases := []struct {
		name   string
		wanted int
		given  [][]int
	}{
		{"case1", 4, [][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}},
		{"case2", 4, [][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}},
		{"case3", 1, [][]int{{1}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, longestIncreasingPath(c.given))
		})
	}
}
