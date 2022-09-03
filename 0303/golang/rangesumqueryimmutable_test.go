package rangesumqueryimmutable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumArray(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)

	assert.Equal(t, 1, obj.SumRange(0, 2))

	assert.Equal(t, -1, obj.SumRange(2, 5))

	assert.Equal(t, -3, obj.SumRange(0, 5))
}
