package shortestworddistance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestDistanceI(t *testing.T) {
	wordDict := []string{"practice", "makes", "perfect", "coding", "makes"}

	result := shortestDistance(wordDict, "coding", "practice")

	assert.Equal(t, 3, result)
}

func TestShortestDistanceII(t *testing.T) {
	wordDict := []string{"practice", "makes", "perfect", "coding", "makes"}

	result := shortestDistance(wordDict, "coding", "makes")

	assert.Equal(t, 1, result)
}
