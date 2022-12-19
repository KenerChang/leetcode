package minimumincrementtomakearrayunique

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinIncrementForUnique(t *testing.T) {
	cases := []struct {
		wanted int
		nums   []int
	}{
		{1, []int{1, 2, 2}},
		{6, []int{3, 2, 1, 2, 1, 7}},
		{1, []int{1, 1, 10}},
		{2, []int{1, 2, 1, 10}},
		{6, []int{1, 2, 3, 1, 2}},
		{0, []int{1, 2, 3}},
		{0, []int{1}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.nums), func(t *testing.T) {
			assert.Equal(t, c.wanted, minIncrementForUnique(c.nums))
		})
	}
}
