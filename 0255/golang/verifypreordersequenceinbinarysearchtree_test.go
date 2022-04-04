package verifypreordersequenceinbinarysearchtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyPreorder(t *testing.T) {
	assert.True(t, verifyPreorder([]int{5, 2, 1, 3, 6}))

	assert.False(t, verifyPreorder([]int{5, 2, 6, 1, 3}))

	assert.False(t, verifyPreorder([]int{5, 2, 6, 1}))

	assert.True(t, verifyPreorder([]int{6}))

	assert.True(t, verifyPreorder([]int{5, 2, 1, 6}))

	assert.False(t, verifyPreorder([]int{1, 3, 4, 2}))
}
