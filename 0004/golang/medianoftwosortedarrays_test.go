package medianoftwosortedarrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	assert.Equal(t, float64(2), findMedianSortedArrays([]int{1, 3}, []int{2}))

	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{1, 3}, []int{2, 4}))

	assert.Equal(t, float64(5.5), findMedianSortedArrays([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}))

	assert.Equal(t, float64(5.5), findMedianSortedArrays([]int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8, 10}))

	assert.Equal(t, float64(5.5), findMedianSortedArrays([]int{1, 2, 3, 9, 10}, []int{4, 5, 6, 7, 8}))

	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{}, []int{2, 3}))

	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{}, []int{1, 2, 3, 4}))

	assert.Equal(t, float64(3), findMedianSortedArrays([]int{2, 2, 4, 4}, []int{2, 2, 4, 4}))

	assert.Equal(t, float64(5), findMedianSortedArrays([]int{1, 3, 5, 7, 9}, []int{1, 3, 5, 7, 9}))

	assert.Equal(t, float64(4), findMedianSortedArrays([]int{1, 3, 5, 7}, []int{1, 3, 5, 7}))
}
