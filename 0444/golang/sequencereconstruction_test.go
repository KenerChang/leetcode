package sequencereconstruction

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequenceReconstruction(t *testing.T) {
	cases := []struct {
		wanted    bool
		nums      []int
		sequences [][]int
	}{
		{false, []int{1, 2, 3}, [][]int{{1, 2}, {1, 3}}},
		{false, []int{1, 2, 3}, [][]int{{1, 2}}},
		{true, []int{1, 2, 3}, [][]int{{1, 2}, {1, 3}, {2, 3}}},
		{false, []int{1, 2, 3, 4}, [][]int{{1, 2}, {1, 3}, {2, 3}}},
		{true, []int{1, 2, 3}, [][]int{{1, 2}, {1, 3}, {1, 2, 3}}},
		{false, []int{1, 2, 3}, [][]int{{2, 1}, {2, 3}}},
		{true, []int{1}, [][]int{{1}}},
		{true, []int{1, 2}, [][]int{{1, 2}}},
		{false, []int{1, 2}, [][]int{{1}, {2}}},
		{true, []int{1, 2}, [][]int{{1}, {2}, {1, 2}}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v_%v", c.nums, c.sequences), func(t *testing.T) {
			assert.Equal(t, c.wanted, sequenceReconstruction(c.nums, c.sequences))
		})
	}
}
