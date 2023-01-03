package longestmountaininarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestMountain(t *testing.T) {
	cases := []struct {
		wanted int
		arr    []int
	}{
		{5, []int{2, 1, 4, 7, 3, 2, 5}},
		{0, []int{2, 2, 2}},
		{3, []int{2, 3, 1}},
		{7, []int{2, 3, 1, 2, 3, 4, 3, 2, 1}},
		{7, []int{2, 3, 1, 2, 3, 4, 3, 2, 1, 1}},
		{4, []int{2, 3, 1, 3, 2, 4, 5, 3}},
		{3, []int{0, 2, 0, 2, 1, 2, 3, 4, 4, 1}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.arr), func(t *testing.T) {
			assert.Equal(t, c.wanted, longestMountain(c.arr))
		})
	}
}
