package peekingiterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPeekIterator(t *testing.T) {
	iter := &Iterator{nums: []int{1, 2, 3}, pos: -1}

	peekIter := Constructor(iter)

	assert.Equal(t, 1, peekIter.next())
	assert.Equal(t, 2, peekIter.peek())
	assert.Equal(t, 2, peekIter.next())
	assert.Equal(t, 3, peekIter.next())

	assert.False(t, peekIter.hasNext())
}
