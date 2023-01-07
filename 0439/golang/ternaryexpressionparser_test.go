package ternaryexpressionparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTernary(t *testing.T) {
	cases := []struct {
		wanted     string
		expression string
	}{
		{"2", "T?2:3"},
		{"4", "F?1:T?4:5"},
		{"F", "T?T?F:5:3"},
		{"T", "F?F:T"},
		{"5", "T?5:T?3:4"},
		{"1", "T?T?T?1:2:3:4"},
		{"4", "F?F?F?1:2:3:4"},
		{"4", "F?T?F?1:2:3:4"},
		{"3", "T?F?T?1:2:3:4"},
	}

	for _, c := range cases {
		t.Run(c.expression, func(t *testing.T) {
			assert.Equal(t, c.wanted, parseTernary(c.expression))
		})
	}
}
