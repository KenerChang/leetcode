package regularexpressionmatching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	assert.False(t, isMatch("aa", "a"))

	assert.True(t, isMatch("aa", "a*"))

	assert.True(t, isMatch("ab", ".*"))

	assert.True(t, isMatch("b", "a*b"))

	assert.False(t, isMatch("bb", "a*b"))

	assert.True(t, isMatch("bcb", ".*b"))

	assert.False(t, isMatch("bcbc", ".*b"))

	assert.True(t, isMatch("mississippi", "mis*is*ip*."))

	assert.True(t, isMatch("a", "."))

	assert.False(t, isMatch("aaa", "aaaa"))

	assert.True(t, isMatch("a", "ab*"))

	assert.True(t, isMatch("a", "b*ab*"))

	assert.True(t, isMatch("ba", "b*ab*"))

	assert.True(t, isMatch("abbabaaaaaaacaa", "a*.*b.a.*c*b*a*c*"))
}
