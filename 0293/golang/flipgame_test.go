package flipgame

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePossibleNextMoves(t *testing.T) {
	cases := []struct {
		currentState string
		wanted       []string
	}{
		{"++++", []string{"--++", "+--+", "++--"}},
		{"+", []string{}},
		{"+--+", []string{}},
		{"+++", []string{"--+", "+--"}},
		{"+++++", []string{"--+++", "+--++", "++--+", "+++--"}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("currentState: %s", c.currentState), func(t *testing.T) {
			assert.ElementsMatch(t, c.wanted, generatePossibleNextMoves(c.currentState))
		})
	}
}
