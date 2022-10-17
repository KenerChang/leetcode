package uniquenumberofoccurrences

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueOccurrences(t *testing.T) {
	cases := []struct {
		wanted bool
		nums   []int
	}{
		{true, []int{1, 2, 2, 1, 1, 3}},
		{false, []int{1, 2}},
		{true, []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.nums), func(t *testing.T) {
			assert.Equal(t, c.wanted, uniqueOccurrences(c.nums))
		})
	}
}
