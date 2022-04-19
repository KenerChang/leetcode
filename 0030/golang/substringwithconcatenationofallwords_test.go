package substringwithconcatenationofallwords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubstring(t *testing.T) {
	assert.ElementsMatch(t, []int{0, 9}, findSubstring("barfoothefoobarman", []string{"foo", "bar"}))

	assert.ElementsMatch(t, []int{}, findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))

	assert.ElementsMatch(t, []int{6, 9, 12}, findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))

	assert.ElementsMatch(t, []int{0, 1, 2}, findSubstring("ffff", []string{"f", "f"}))

	assert.ElementsMatch(t, []int{}, findSubstring("f", []string{"f", "f"}))

	assert.ElementsMatch(t, []int{0}, findSubstring("ff", []string{"f", "f"}))

	assert.ElementsMatch(t, []int{0}, findSubstring("f", []string{"f"}))

	assert.ElementsMatch(t, []int{0, 1}, findSubstring("ff", []string{"f"}))

	assert.ElementsMatch(t, []int{0}, findSubstring("wordgoodgoodgoodbestword", []string{"wordgoodgoodgoodbestword"}))

	assert.ElementsMatch(t, []int{}, findSubstring("wordgoodgoodgoodbestword", []string{"wordgoodgoodgoodbestwordccccc"}))
}
