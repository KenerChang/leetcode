package outputcontestmatches

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindContestMatch(t *testing.T) {
	cases := []struct {
		wanted string
		n      int
	}{
		{"((1,4),(2,3))", 4},
		{"(((1,8),(4,5)),((2,7),(3,6)))", 8},
		{"((((1,16),(8,9)),((4,13),(5,12))),(((2,15),(7,10)),((3,14),(6,11))))", 16},
		{"(1,2)", 2},
		{"(((((1,32),(16,17)),((8,25),(9,24))),(((4,29),(13,20)),((5,28),(12,21)))),((((2,31),(15,18)),((7,26),(10,23))),(((3,30),(14,19)),((6,27),(11,22)))))", 32},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.n), func(t *testing.T) {
			assert.Equal(t, c.wanted, findContestMatch(c.n))
		})
	}
}
