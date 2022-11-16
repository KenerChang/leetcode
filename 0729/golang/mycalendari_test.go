package mycalendari

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyCalendar(t *testing.T) {
	c := Constructor()

	assert.True(t, c.Book(10, 20))
	assert.False(t, c.Book(15, 25))
	assert.True(t, c.Book(20, 30))
}

func TestMyCalendarII(t *testing.T) {
	c := Constructor()

	assert.True(t, c.Book(10, 20))
	assert.True(t, c.Book(20, 30))
	assert.False(t, c.Book(15, 25))
}

func TestMyCalendarIII(t *testing.T) {
	c := Constructor()

	assert.True(t, c.Book(10, 20))
	assert.True(t, c.Book(30, 40))
	assert.True(t, c.Book(20, 25))
}

func TestMyCalendarIV(t *testing.T) {
	c := Constructor()

	assert.True(t, c.Book(47, 50))
	assert.True(t, c.Book(33, 41))
	assert.False(t, c.Book(39, 45))
	assert.False(t, c.Book(33, 42))
	assert.True(t, c.Book(25, 32))
	assert.False(t, c.Book(26, 35))
	assert.True(t, c.Book(19, 25))
	assert.True(t, c.Book(3, 8))
	assert.True(t, c.Book(8, 13))
	assert.False(t, c.Book(18, 27))
}

func TestMyCalendarV(t *testing.T) {
	c := Constructor()

	assert.True(t, c.Book(37, 50))
	assert.False(t, c.Book(33, 50))
	assert.True(t, c.Book(4, 17))
	assert.False(t, c.Book(35, 48))
	assert.False(t, c.Book(8, 25))
}
