package stringcompression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompressI(t *testing.T) {
	input := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	result := compress(input)

	assert.Equal(t, "a2b2c3", string(input[0:result]))
	assert.Equal(t, 6, result)
}

func TestCompressII(t *testing.T) {
	input := []byte{'a'}
	result := compress(input)

	assert.Equal(t, "a", string(input[0:result]))
	assert.Equal(t, 1, result)
}

func TestCompressIII(t *testing.T) {
	input := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	result := compress(input)

	assert.Equal(t, "ab12", string(input[0:result]))
	assert.Equal(t, 4, result)
}

func TestCompressIV(t *testing.T) {
	input := []byte{'a', 'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	result := compress(input)

	assert.Equal(t, "a2b12", string(input[0:result]))
	assert.Equal(t, 5, result)
}

func TestCompressV(t *testing.T) {
	input := []byte{'a', 'a', 'b', 'c'}
	result := compress(input)

	assert.Equal(t, "a2bc", string(input[0:result]))
	assert.Equal(t, 4, result)
}
