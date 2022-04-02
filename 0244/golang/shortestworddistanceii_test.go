package shortestworddistanceii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordDistanceI(t *testing.T) {
	wd := Constructor([]string{"practice", "makes", "perfect", "coding", "makes"})

	assert.Equal(t, 3, wd.Shortest("coding", "practice"))
	assert.Equal(t, 1, wd.Shortest("makes", "coding"))
}
