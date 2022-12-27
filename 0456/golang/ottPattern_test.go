package ottPattern

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind132pattern(t *testing.T) {
	cases := []struct {
		wanted bool
		nums   []int
	}{
		{false, []int{1, 2, 3, 4}},
		{true, []int{3, 1, 4, 2}},
		{true, []int{-1, 3, 2, 0}},
		{false, []int{5, 4, 1, 2, 3}},
		{true, []int{5, 1, 4, 2, 3}},
		{true, []int{5, 4, 1, 3, 2}},
		{false, []int{1}},
		{false, []int{1, 3}},
		{true, []int{1, 3, 2}},
		{false, []int{1, 2, 3}},
		{false, []int{1, 0, 1, -4, -3}},
		{true, []int{1, 0, 1, -4, -2, -3}},
		{true, []int{1000, 10, 15, -4, -2, -3, 11}},
		{true, []int{1000, 10, 15, -4, -2, 11}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.nums), func(t *testing.T) {
			assert.Equal(t, c.wanted, find132pattern(c.nums))
		})
	}
}
