package rotatestring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateString(t *testing.T) {
	tests := []struct {
		s, goal string
		want    bool
	}{
		{"abcde", "cdeab", true},
		{"abcde", "abced", false},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			assert.Equal(t, test.want, rotateString(test.s, test.goal))
		})
	}
}
