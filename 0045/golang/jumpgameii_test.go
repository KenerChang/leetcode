package jumpgameii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	cases := []struct {
		nums   []int
		wanted int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{4, 3, 1, 1, 4}, 1},
		{[]int{1}, 0},
		{[]int{0}, 0},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.nums), func(t *testing.T) {
			assert.Equal(t, c.wanted, jump(c.nums))
		})
	}
}
