package aliendictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlienOrder(t *testing.T) {
	assert.Equal(t, "wertf", alienOrder([]string{"wrt", "wrf", "er", "ett", "rftt"}))

	assert.Equal(t, "zx", alienOrder([]string{"z", "x"}))

	assert.Equal(t, "", alienOrder([]string{"z", "x", "z"}))

	assert.Equal(t, "z", alienOrder([]string{"z", "z"}))

	assert.Equal(t, "zxy", alienOrder([]string{"zx", "zy"}))
	assert.Equal(t, "z", alienOrder([]string{"z", "z", "z"}))

	assert.Equal(t, "z", alienOrder([]string{"z"}))

	assert.Equal(t, "wx", alienOrder([]string{"wx"}))

	assert.Equal(t, "abcd", alienOrder([]string{"ab", "adc"}))

	assert.Equal(t, "", alienOrder([]string{"abc", "ab"}))

	assert.Equal(t, "abc", alienOrder([]string{"ab", "abc"}))

	assert.Equal(t, "xbca", alienOrder([]string{"x", "ab", "abc"}))

	assert.Equal(t, "", alienOrder([]string{"x", "abc", "ab"}))

	assert.Equal(t, "", alienOrder([]string{"z", "x", "a", "zb", "zx"}))

	assert.Equal(t, "ab", alienOrder([]string{"aba"}))
}
