package flatten2dvector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVector2DI(t *testing.T) {
	v := Constructor([][]int{
		{1, 2},
		{3},
		{4},
	})

	assert.Equal(t, 1, v.Next())
	assert.Equal(t, 2, v.Next())
	assert.Equal(t, 3, v.Next())
	assert.True(t, v.HasNext())
	assert.True(t, v.HasNext())
	assert.Equal(t, 4, v.Next())
	assert.False(t, v.HasNext())
}

func TestVector2DII(t *testing.T) {
	v := Constructor([][]int{
		{},
		{3},
		{4},
	})

	assert.Equal(t, 3, v.Next())
	assert.True(t, v.HasNext())
	assert.True(t, v.HasNext())
	assert.Equal(t, 4, v.Next())
	assert.False(t, v.HasNext())
}

func TestVector2DIII(t *testing.T) {
	v := Constructor([][]int{})

	assert.False(t, v.HasNext())
}
