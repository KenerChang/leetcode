package palindromepermutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPermutePalindrome(t *testing.T) {
	assert.False(t, canPermutePalindrome("code"))

	assert.True(t, canPermutePalindrome("aab"))

	assert.True(t, canPermutePalindrome("carerac"))

	assert.True(t, canPermutePalindrome("a"))

	assert.False(t, canPermutePalindrome("ab"))

	assert.False(t, canPermutePalindrome("abc"))
}
