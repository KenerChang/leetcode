package bullsandcows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHint(t *testing.T) {
	tests := []struct {
		secret string
		guess  string
		want   string
	}{
		{"1807", "7810", "1A3B"},
		{"1123", "0111", "1A1B"},
		{"1122", "1222", "3A0B"},
		{"1122", "2211", "0A4B"},
		{"1111", "1111", "4A0B"},
		{"1234", "5678", "0A0B"},
	}
	for _, tt := range tests {
		t.Run(tt.secret, func(t *testing.T) {
			assert.Equal(t, tt.want, getHint(tt.secret, tt.guess))
		})
	}
}
