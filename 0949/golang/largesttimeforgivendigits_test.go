package largesttimeforgivendigits

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestTimeFromDigits(t *testing.T) {
	cases := []struct {
		wanted string
		arr    []int
	}{
		{"23:41", []int{1, 2, 3, 4}},
		{"", []int{5, 5, 5, 5}},
		{"11:11", []int{1, 1, 1, 1}},
		{"00:00", []int{0, 0, 0, 0}},
		{"22:11", []int{2, 2, 1, 1}},
		{"22:11", []int{1, 1, 2, 2}},
		{"04:00", []int{4, 0, 0, 0}},
		{"04:00", []int{4, 0, 0, 0}},
		{"19:06", []int{1, 9, 6, 0}},
		{"06:26", []int{2, 0, 6, 6}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.arr), func(t *testing.T) {
			assert.Equal(t, c.wanted, largestTimeFromDigits(c.arr))
		})
	}
}
