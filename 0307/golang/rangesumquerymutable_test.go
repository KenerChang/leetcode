package rangesumquerymutable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumArray(t *testing.T) {
	numArray := Constructor([]int{1, 3, 5})

	assert.Equal(t, numArray.SumRange(0, 2), 9)

	numArray.Update(1, 2)

	assert.Equal(t, numArray.SumRange(0, 2), 8)
}

func TestNumArrayII(t *testing.T) {
	numArray := Constructor([]int{7, 2, 7, 2, 0})

	numArray.Update(4, 6)

	numArray.Update(0, 2)

	numArray.Update(0, 9)

	assert.Equal(t, numArray.SumRange(4, 4), 6)

	numArray.Update(3, 8)

	assert.Equal(t, numArray.SumRange(0, 4), 32)

	numArray.Update(4, 1)

	assert.Equal(t, numArray.SumRange(0, 3), 26)

	assert.Equal(t, numArray.SumRange(0, 4), 27)

	numArray.Update(0, 4)
}
