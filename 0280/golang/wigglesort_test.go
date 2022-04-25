package wigglesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isWiggleSort(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if i%2 == 0 {
			if nums[i] > nums[i+1] {
				return false
			}
		} else {
			if nums[i] < nums[i+1] {
				return false
			}
		}
	}

	return true
}

func TestWiggleSort(t *testing.T) {
	input := []int{3, 5, 2, 1, 6, 4}
	wiggleSort(input)

	assert.True(t, isWiggleSort(input))
}

func TestWiggleSortII(t *testing.T) {
	input := []int{6, 6, 5, 6, 3, 8}
	wiggleSort(input)

	assert.True(t, isWiggleSort(input))
}

func TestWiggleSortIII(t *testing.T) {
	input := []int{2, 1}
	wiggleSort(input)

	assert.True(t, isWiggleSort(input))
}

func TestWiggleSortIV(t *testing.T) {
	input := []int{1, 2}
	wiggleSort(input)

	assert.True(t, isWiggleSort(input))
}
