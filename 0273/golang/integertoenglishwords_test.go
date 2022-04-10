package integertoenglishwords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberToWords(t *testing.T) {
	assert.Equal(t, "One Hundred Twenty Three", numberToWords(123))

	assert.Equal(t, "One Hundred Twenty Three Thousand Four Hundred Fifty Six", numberToWords(123456))
	assert.Equal(t, "Twelve Million Three Hundred Forty Five Thousand Six Hundred Seventy Eight", numberToWords(12345678))

	assert.Equal(t, "One Hundred Twenty Three Million Four Hundred Fifty Six Thousand Seven Hundred Eighty Nine", numberToWords(123456789))

	assert.Equal(t, "Two Billion One Hundred Twenty Three Million Four Hundred Fifty Six Thousand Seven Hundred Eighty Nine", numberToWords(2123456789))

	assert.Equal(t, "Twelve Thousand Three Hundred Forty Five", numberToWords(12345))

	assert.Equal(t, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven", numberToWords(1234567))

	assert.Equal(t, "Zero", numberToWords(0))

	assert.Equal(t, "Five Hundred One", numberToWords(501))

	assert.Equal(t, "Fifteen", numberToWords(15))

	assert.Equal(t, "Twenty Five", numberToWords(25))

	assert.Equal(t, "Fifty Four", numberToWords(54))

	assert.Equal(t, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven", numberToWords(1234567))

	assert.Equal(t, "Ten", numberToWords(10))

	assert.Equal(t, "One Hundred", numberToWords(100))

	assert.Equal(t, "One Thousand", numberToWords(1000))
	assert.Equal(t, "Ten Thousand", numberToWords(10000))
	assert.Equal(t, "One Million", numberToWords(1000000))
	assert.Equal(t, "One Hundred Million", numberToWords(100000000))
	assert.Equal(t, "One Billion", numberToWords(1000000000))

	assert.Equal(t, "One Thousand One", numberToWords(1001))

	assert.Equal(t, "Fifty Thousand Eight Hundred Sixty Eight", numberToWords(50868))
}
