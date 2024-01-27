package uniquewordabbreviation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidWordAbbr(t *testing.T) {
	validator := Constructor([]string{"deer", "door", "cake", "card", "cane", "cxxe"})

	assert.False(t, validator.IsUnique("dear"))
	assert.True(t, validator.IsUnique("cart"))
	assert.False(t, validator.IsUnique("cane"))
	assert.True(t, validator.IsUnique("make"))
	assert.False(t, validator.IsUnique("cake"))
}

func TestValidWordAbbrII(t *testing.T) {
	validator := Constructor([]string{"deer", "door", "cake", "card", "cake"})

	assert.False(t, validator.IsUnique("dear"))
	assert.True(t, validator.IsUnique("cart"))
	assert.False(t, validator.IsUnique("cane"))
	assert.True(t, validator.IsUnique("make"))
	assert.True(t, validator.IsUnique("cake"))
}
