package movingaveragefromdatastream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingAverageI(t *testing.T) {
	m := Constructor(3)

	assert.Equal(t, 1.0, m.Next(1))

	assert.Equal(t, 5.5, m.Next(10))

	assert.Equal(t, float64(1+10+3)/3, m.Next(3))

	assert.Equal(t, 6.0, m.Next(5))
}
