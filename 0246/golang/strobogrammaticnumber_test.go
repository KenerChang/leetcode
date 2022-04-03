package strobogrammaticnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStrobogrammatic(t *testing.T) {
	assert.True(t, isStrobogrammatic("69"))

	assert.True(t, isStrobogrammatic("88"))

	assert.False(t, isStrobogrammatic("962"))

	assert.False(t, isStrobogrammatic("2"))

	assert.True(t, isStrobogrammatic("888"))

	assert.True(t, isStrobogrammatic("0"))

	assert.True(t, isStrobogrammatic("101"))
}
