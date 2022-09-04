package minimumwindowsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinWindow(t *testing.T) {
	assert.Equal(t, "BANC", minWindow("ADOBECODEBANC", "ABC"))

	assert.Equal(t, "BECODEBA", minWindow("ADOBECODEBANC", "ABBC"))

	assert.Equal(t, "a", minWindow("a", "a"))

	assert.Equal(t, "", minWindow("a", "aa"))

}
