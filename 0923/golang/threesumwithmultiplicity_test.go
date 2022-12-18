package threesumwithmultiplicity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSumMulti(t *testing.T) {
	cases := []struct {
		wanted int
		arr    []int
		target int
	}{
		{20, []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8},
		{12, []int{1, 1, 2, 2, 2, 2}, 5},
		{1, []int{2, 1, 3}, 6},
		{3, []int{1, 1, 3, 2, 2}, 5},
		{1, []int{1, 5, 6, 7, 1, 1}, 3},
		{0, []int{5, 5, 5}, 1},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("arr:%v,target:%d", c.arr, c.target), func(t *testing.T) {
			assert.Equal(t, c.wanted, threeSumMulti(c.arr, c.target))
		})
	}
}
