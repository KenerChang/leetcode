package nextgreaterelementi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextGreaterElement(t *testing.T) {
	cases := []struct {
		wanted []int
		nums1  []int
		nums2  []int
	}{
		{[]int{-1, 3, -1}, []int{4, 1, 2}, []int{1, 3, 4, 2}},
		{[]int{3, -1}, []int{2, 4}, []int{1, 2, 3, 4}},
		{[]int{-1}, []int{1}, []int{1}},
		{[]int{-1, -1, -1, -1, -1}, []int{5, 4, 3, 2, 1}, []int{5, 4, 3, 2, 1}},
		{[]int{2, 3, 4, 5, -1}, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{-1, -1, -1, -1, -1}, []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{-1, 5, 4, 3, 2}, []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("nums1:%v,nums2:%v", c.nums1, c.nums2), func(t *testing.T) {
			assert.Equal(t, c.wanted, nextGreaterElement(c.nums1, c.nums2))
		})
	}
}
