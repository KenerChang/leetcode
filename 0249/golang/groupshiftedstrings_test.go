package groupshiftedstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupStrings(t *testing.T) {
	input := []string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}

	target := [][]string{
		{"acef"}, {"a", "z"}, {"abc", "bcd", "xyz"}, {"az", "ba"},
	}

	assert.ElementsMatch(t, target, groupStrings(input))

	input = []string{"a"}

	target = [][]string{
		{"a"},
	}

	assert.Equal(t, target, groupStrings(input))
}
