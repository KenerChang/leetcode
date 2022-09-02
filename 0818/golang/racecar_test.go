package racecar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRaceCar(t *testing.T) {
	assert.Equal(t, 0, racecar(0))

	assert.Equal(t, 4, racecar(2))

	assert.Equal(t, 2, racecar(3))

	assert.Equal(t, 7, racecar(5))

	assert.Equal(t, 5, racecar(4))
	assert.Equal(t, 5, racecar(6))

	assert.Equal(t, 6, racecar(8))
}
