package wordladderii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLadders(t *testing.T) {
	target := [][]string{
		{"hit", "hot", "dot", "dog", "cog"},
		{"hit", "hot", "lot", "log", "cog"},
	}

	result := findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})

	assert.ElementsMatch(t, target, result)
}

func TestFindLaddersII(t *testing.T) {
	target := [][]string{}

	result := findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"})

	assert.ElementsMatch(t, target, result)
}

func TestFindLaddersIII(t *testing.T) {
	target := [][]string{
		{"hit", "hot", "dot", "dog"},
	}

	result := findLadders("hit", "dog", []string{"hot", "dot", "dog", "lot", "log", "cog"})

	assert.ElementsMatch(t, target, result)
}

func TestFindLaddersIV(t *testing.T) {
	target := [][]string{
		{"hit", "hot"},
	}

	result := findLadders("hit", "hot", []string{"hot", "dit", "dot"})

	assert.ElementsMatch(t, target, result)
}

func TestFindLaddersV(t *testing.T) {
	target := [][]string{}

	result := findLadders("hit", "hot", []string{"dit", "dot"})

	assert.ElementsMatch(t, target, result)
}
