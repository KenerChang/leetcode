package twosum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		wanted []int
		nums   []int
		target int
	}{
		{[]int{0, 1}, []int{2, 7, 11, 15}, 9},
		{[]int{1, 2}, []int{3, 2, 4}, 6},
		{[]int{0, 1}, []int{3, 3}, 6},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("nums:%v,target:%d", c.nums, c.target), func(t *testing.T) {
			assert.Equal(t, c.wanted, twoSum(c.nums, c.target))
		})
	}
}
