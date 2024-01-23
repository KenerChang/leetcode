package insertinterval

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	cases := []struct {
		intervals   [][]int
		newInterval []int
		wanted      [][]int
	}{
		{[][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
		{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{[][]int{{2, 3}, {6, 9}}, []int{0, 1}, [][]int{{0, 1}, {2, 3}, {6, 9}}},
		{[][]int{{2, 3}, {6, 9}}, []int{10, 11}, [][]int{{2, 3}, {6, 9}, {10, 11}}},
		{[][]int{{1, 3}, {6, 9}}, []int{1, 3}, [][]int{{1, 3}, {6, 9}}},
		{[][]int{{1, 3}, {6, 9}}, []int{0, 4}, [][]int{{0, 4}, {6, 9}}},
		{[][]int{{1, 3}, {6, 9}}, []int{0, 10}, [][]int{{0, 10}}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.intervals), func(t *testing.T) {
			assert.Equal(t, c.wanted, insert(c.intervals, c.newInterval))
		})
	}
}
