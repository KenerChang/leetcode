package reconstructitinerary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindItinerary(t *testing.T) {
	cases := []struct {
		name    string
		tickets [][]string
		wanted  []string
	}{
		{"case 1", [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}, []string{"JFK", "MUC", "LHR", "SFO", "SJC"}},
		{"case 2", [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}, []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, findItinerary(c.tickets))
		})
	}
}
