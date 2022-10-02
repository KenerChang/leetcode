package maxstack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxStack(t *testing.T) {
	s := Constructor()
	s.Push(5)
	s.Push(1)
	s.Push(5)
	assert.Equal(t, 5, s.Top())
	assert.Equal(t, 5, s.PopMax())
	assert.Equal(t, 1, s.Top())
	assert.Equal(t, 5, s.PeekMax())
	assert.Equal(t, 1, s.Pop())
	assert.Equal(t, 5, s.Top())
}
