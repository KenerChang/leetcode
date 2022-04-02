package shortestworddistanceiii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestWordDistanceI(t *testing.T) {
	wordsDict := []string{"practice", "makes", "perfect", "coding", "makes"}

	assert.Equal(t, 1, shortestWordDistance(wordsDict, "makes", "coding"))
}

func TestShortestWordDistanceII(t *testing.T) {
	wordsDict := []string{"practice", "makes", "perfect", "coding", "makes"}

	assert.Equal(t, 3, shortestWordDistance(wordsDict, "makes", "makes"))
}
