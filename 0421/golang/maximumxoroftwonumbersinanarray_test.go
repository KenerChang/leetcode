package maximumxoroftwonumbersinanarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaximumXOR(t *testing.T) {
	cases := []struct {
		wanted int
		nums   []int
	}{
		{28, []int{3, 10, 5, 25, 2, 8}},
		{127, []int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}},
		{1, []int{1, 0}},
		{0, []int{1}},
		{0, []int{3}},
		{0, []int{0}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.nums), func(t *testing.T) {
			assert.Equal(t, c.wanted, findMaximumXOR(c.nums))
		})
	}
}
