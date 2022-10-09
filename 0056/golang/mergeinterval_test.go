package mergeinterval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	cases := []struct {
		name      string
		wanted    [][]int
		intervals [][]int
	}{
		{"case 1", [][]int{{1, 6}, {8, 10}, {15, 18}}, [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
		{"case 2", [][]int{{1, 5}}, [][]int{{1, 4}, {4, 5}}},
		{"case 3", [][]int{{1, 4}}, [][]int{{1, 4}, {2, 3}}},
		{"case 4", [][]int{{0, 4}}, [][]int{{1, 4}, {0, 4}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, merge(c.intervals))
		})
	}
}
