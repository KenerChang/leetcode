package binarysearch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		wanted int
		nums   []int
		target int
	}{
		{4, []int{-1, 0, 3, 5, 9, 12}, 9},
		{-1, []int{-1, 0, 3, 5, 9, 12}, 2},
		{5, []int{-1, 0, 3, 5, 9, 12}, 12},
		{0, []int{-1, 0, 3, 5, 9, 12}, -1},
		{-1, []int{-1, 0, 3, 5, 9, 12}, 13},
		{-1, []int{-1, 0, 3, 5, 9, 12}, -5},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.nums), func(t *testing.T) {
			assert.Equal(t, c.wanted, search(c.nums, c.target))
		})
	}
}
