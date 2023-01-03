package similarrgbcolor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimilarRGB(t *testing.T) {
	cases := []struct {
		wanted string
		color  string
	}{
		{"#11ee66", "#09f166"},
		{"#5544dd", "#4e3fe1"},
		{"#112233", "#112233"},
		{"#ffffff", "#ffffff"},
		{"#000000", "#000000"},
	}

	for _, c := range cases {
		t.Run(c.color, func(t *testing.T) {
			assert.Equal(t, c.wanted, similarRGB(c.color))
		})
	}
}
