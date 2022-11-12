package maximumlengthofrepeatedsubarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLength(t *testing.T) {
	cases := []struct {
		wanted int
		nums1  []int
		nums2  []int
	}{
		{3, []int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}},
		{5, []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}},
		{2, []int{0, 1, 1, 1, 1}, []int{1, 0, 1, 0, 1}},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, c.wanted, findLength(c.nums1, c.nums2))
		})
	}
}
