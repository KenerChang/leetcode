package russiandollenvelopes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxEnvelopes(t *testing.T) {
	cases := []struct {
		wanted    int
		envelopes [][]int
	}{
		{3, [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}},
		{1, [][]int{{1, 1}, {1, 1}, {1, 1}}},
		{3, [][]int{{2, 3}, {3, 4}, {4, 5}}},
		{1, [][]int{{1, 1}}},
		{5, [][]int{{2, 100}, {3, 200}, {4, 300}, {5, 500}, {5, 400}, {5, 250}, {6, 370}, {6, 360}, {7, 380}}},
		{5, [][]int{{15, 8}, {2, 20}, {2, 14}, {4, 17}, {8, 19}, {8, 9}, {5, 7}, {11, 19}, {8, 11}, {13, 11}, {2, 13}, {11, 19}, {8, 11}, {13, 11}, {2, 13}, {11, 19}, {16, 1}, {18, 13}, {14, 17}, {18, 19}}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.envelopes), func(t *testing.T) {
			assert.Equal(t, c.wanted, maxEnvelopes(c.envelopes))
		})
	}
}
