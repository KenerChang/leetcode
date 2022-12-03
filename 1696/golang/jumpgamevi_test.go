package jumpgamevi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxResult(t *testing.T) {
	cases := []struct {
		wanted int
		nums   []int
		k      int
	}{
		{7, []int{1, -1, -2, 4, -7, 3}, 2},
		{17, []int{10, -5, -2, 4, 0, 3}, 3},
		{0, []int{1, -5, -20, 4, -1, 3, -6, -3}, 2},
		{26, []int{5, 6, 7, 8}, 3},
		{13, []int{5, -6, -7, 8}, 3},
		{13, []int{5, 8}, 3},
		{198, []int{100, -1, -100, -1, 100}, 2},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v_%d", c.nums, c.k), func(t *testing.T) {
			assert.Equal(t, c.wanted, maxResult(c.nums, c.k))
		})
	}
}
