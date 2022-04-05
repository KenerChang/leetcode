package uglynumberii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthUglyNumber(t *testing.T) {
	assert.Equal(t, 12, nthUglyNumber(10))

	assert.Equal(t, 1, nthUglyNumber(1))
	assert.Equal(t, 2, nthUglyNumber(2))
	assert.Equal(t, 3, nthUglyNumber(3))

	assert.Equal(t, 6, nthUglyNumber(6))
	assert.Equal(t, 8, nthUglyNumber(7))
	assert.Equal(t, 9, nthUglyNumber(8))
	assert.Equal(t, 10, nthUglyNumber(9))
	assert.Equal(t, 12, nthUglyNumber(10))
	assert.Equal(t, 24, nthUglyNumber(15))
	assert.Equal(t, 2123366400, nthUglyNumber(1690))
}
