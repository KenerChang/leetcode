package palindromepermutationii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePalindromes(t *testing.T) {
	target := []string{"abba", "baab"}

	assert.ElementsMatch(t, target, generatePalindromes("abba"))

	target = []string{}

	assert.ElementsMatch(t, target, generatePalindromes("abc"))

	target = []string{"a"}

	assert.ElementsMatch(t, target, generatePalindromes("a"))

	target = []string{"aba"}
	assert.ElementsMatch(t, target, generatePalindromes("aab"))

}
