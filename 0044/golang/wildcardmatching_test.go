package wildcardmatching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWildcardMatch(t *testing.T) {
	assert.False(t, isMatch("a", "aa"))

	assert.True(t, isMatch("a", "*"))

	assert.False(t, isMatch("cb", "?a"))

	assert.False(t, isMatch("aa", "*b"))

	assert.True(t, isMatch("aab", "*b"))

	assert.True(t, isMatch("ab", "ab*"))

	assert.False(t, isMatch("ab", "ab?"))

	assert.True(t, isMatch("ab", "ab**"))

	assert.True(t, isMatch("ab", "*ab"))

	assert.True(t, isMatch("ab", "**ab"))

	assert.False(t, isMatch("ab", "?ab"))

	assert.True(t, isMatch("cab", "?ab"))

	assert.True(t, isMatch("cab", "?*ab"))

	assert.True(t, isMatch(
		"aaaabaaaabbbbaabbbaabbaababbabbaaaababaaabbbbbbaabbbabababbaaabaabaaaaaabbaabbbbaababbababaabbbaababbbba",
		"*****b*aba***babaa*bbaba***a*aaba*b*aa**a*b**ba***a*a*",
	),
	)

}
