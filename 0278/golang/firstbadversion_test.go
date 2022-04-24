package firstbadversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isBadVersionFactory(badVersion int) func(int) bool {
	return func(version int) bool {
		return version >= badVersion
	}
}

func TestFirstBadVersionI(t *testing.T) {
	isBadVersion = isBadVersionFactory(4)

	assert.Equal(t, 4, firstBadVersion(4))
}

func TestFirstBadVersionII(t *testing.T) {
	isBadVersion = isBadVersionFactory(1)

	assert.Equal(t, 1, firstBadVersion(1))
}

func TestFirstBadVersionIII(t *testing.T) {
	isBadVersion = isBadVersionFactory(2)

	assert.Equal(t, 2, firstBadVersion(5))
}
