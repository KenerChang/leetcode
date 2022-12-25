package designphonedirectory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhoneDirectoryI(t *testing.T) {
	d := Constructor(3)

	assert.Equal(t, 0, d.Get())
	assert.Equal(t, 1, d.Get())
	assert.True(t, d.Check(2))
	assert.Equal(t, 2, d.Get())

	d.Release(2)

	assert.True(t, d.Check(2))
}
