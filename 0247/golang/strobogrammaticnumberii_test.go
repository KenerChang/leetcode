package strobogrammaticnumberii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindStrobogrammatic(t *testing.T) {
	target := []string{"11", "69", "88", "96"}

	assert.Equal(t, target, findStrobogrammatic(2))

	target = []string{"0", "1", "8"}

	assert.Equal(t, target, findStrobogrammatic(1))

	target = []string{"101", "111", "181", "609", "619", "689", "808", "818", "888", "906", "916", "986"}

	assert.Equal(t, target, findStrobogrammatic(3))
}
