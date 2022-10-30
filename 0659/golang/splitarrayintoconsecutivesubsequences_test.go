package splitarrayintoconsecutivesubsequences

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPossible(t *testing.T) {
	cases := []struct {
		wanted bool
		nums   []int
	}{
		{true, []int{1, 2, 3, 3, 4, 5}},
		{true, []int{1, 2, 3, 3, 4, 4, 5, 5}},
		{false, []int{1, 2, 3, 4, 4, 5}},
		{true, []int{1, 2, 3, 3, 4, 5, 5, 6, 7}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.nums), func(t *testing.T) {
			assert.Equal(t, c.wanted, isPossible(c.nums))
		})
	}
}
