package twosumiii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSumI(t *testing.T) {

	twoSum := Constructor()
	twoSum.Add(1)                  // [] --> [1]
	twoSum.Add(3)                  // [1] --> [1,3]
	twoSum.Add(5)                  // [1,3] --> [1,3,5]
	assert.True(t, twoSum.Find(4)) // 1 + 3 = 4, return true
	assert.False(t, twoSum.Find(7))
}

func TestTwoSumII(t *testing.T) {

	twoSum := Constructor()
	twoSum.Add(0) // [] --> [0]
	assert.False(t, twoSum.Find(0))
}

func TestTwoSumIII(t *testing.T) {

	twoSum := Constructor()
	twoSum.Add(0) // [] --> [0]
	twoSum.Add(0) // [] --> [0, 0]
	assert.True(t, twoSum.Find(0))
}
