package evaluatedivision

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcEquation(t *testing.T) {
	cases := []struct {
		name      string
		wanted    []float64
		equations [][]string
		values    []float64
		queries   [][]string
	}{
		// {"case 1", []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000}, [][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}},
		// {"case 2", []float64{3.75000, 0.40000, 5.00000, 0.20000}, [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}, []float64{1.5, 2.5, 5.0}, [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}},
		// {"case 3", []float64{0.50000, 2.00000, -1.00000, -1.00000}, [][]string{{"a", "b"}}, []float64{0.5}, [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}}},
		{
			"case 4",
			[]float64{0.66667, 0.33333, -1.00000, 1.00000, 1.00000, 0.04464},
			[][]string{{"a", "b"}, {"c", "b"}, {"d", "b"}, {"w", "x"}, {"y", "x"}, {"z", "x"}, {"w", "d"}},
			[]float64{2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0},
			[][]string{{"a", "c"}, {"b", "c"}, {"a", "e"}, {"a", "a"}, {"x", "x"}, {"a", "z"}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, calcEquation(c.equations, c.values, c.queries))
		})
	}
}
