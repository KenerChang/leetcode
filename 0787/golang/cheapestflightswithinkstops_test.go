package cheapestflightswithinkstops

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCheapestPrice(t *testing.T) {
	cases := []struct {
		name    string
		wanted  int
		n       int
		flights [][]int
		src     int
		dst     int
		k       int
	}{
		{"case 1", 700, 4, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}, {1, 3, 600}, {2, 3, 200}}, 0, 3, 1},
		{"case 2", 200, 3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1},
		{"case 3", 500, 3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0},
		{"case 4", 169, 18, [][]int{{16, 1, 81}, {15, 13, 47}, {1, 0, 24}, {5, 10, 21}, {7, 1, 72}, {0, 4, 88}, {16, 4, 39}, {9, 3, 25}, {10, 11, 28}, {13, 8, 93}, {10, 3, 62}, {14, 0, 38}, {3, 10, 58}, {3, 12, 46}, {3, 8, 2}, {10, 16, 27}, {6, 9, 90}, {14, 8, 6}, {0, 13, 31}, {6, 4, 65}, {14, 17, 29}, {13, 17, 64}, {12, 5, 26}, {12, 1, 9}, {12, 15, 79}, {16, 11, 79}, {16, 15, 17}, {4, 0, 21}, {15, 10, 75}, {3, 17, 23}, {8, 5, 55}, {9, 4, 19}, {0, 10, 83}, {3, 7, 17}, {0, 12, 31}, {11, 5, 34}, {17, 14, 98}, {11, 14, 85}, {16, 7, 48}, {12, 6, 86}, {5, 17, 72}, {4, 12, 5}, {12, 10, 23}, {3, 2, 31}, {12, 7, 5}, {6, 13, 30}, {6, 7, 88}, {2, 17, 88}, {6, 8, 98}, {0, 7, 69}, {10, 15, 13}, {16, 14, 24}, {1, 17, 24}, {13, 9, 82}, {13, 6, 67}, {15, 11, 72}, {12, 0, 83}, {1, 4, 37}, {12, 9, 36}, {9, 17, 81}, {9, 15, 62}, {8, 15, 71}, {10, 12, 25}, {7, 6, 23}, {16, 5, 76}, {7, 17, 4}, {3, 11, 82}, {2, 11, 71}, {8, 4, 11}, {14, 10, 51}, {8, 10, 51}, {4, 1, 57}, {6, 16, 68}, {3, 9, 100}, {1, 14, 26}, {10, 7, 14}, {8, 17, 24}, {1, 11, 10}, {2, 9, 85}, {9, 6, 49}, {11, 4, 95}}, 7, 2, 6},
		{"case 5", 169, 15, [][]int{{10, 14, 43}, {1, 12, 62}, {4, 2, 62}, {14, 10, 49}, {9, 5, 29}, {13, 7, 53}, {4, 12, 90}, {14, 9, 38}, {11, 2, 64}, {2, 13, 92}, {11, 5, 42}, {10, 1, 89}, {14, 0, 32}, {9, 4, 81}, {3, 6, 97}, {7, 13, 35}, {11, 9, 63}, {5, 7, 82}, {13, 6, 57}, {4, 5, 100}, {2, 9, 34}, {11, 13, 1}, {14, 8, 1}, {12, 10, 42}, {2, 4, 41}, {0, 6, 55}, {5, 12, 1}, {13, 3, 67}, {3, 13, 36}, {3, 12, 73}, {7, 5, 72}, {5, 6, 100}, {7, 6, 52}, {4, 7, 43}, {6, 3, 67}, {3, 1, 66}, {8, 12, 30}, {8, 3, 42}, {9, 3, 57}, {12, 6, 31}, {2, 7, 10}, {14, 4, 91}, {2, 3, 29}, {8, 9, 29}, {2, 11, 65}, {3, 8, 49}, {6, 14, 22}, {4, 6, 38}, {13, 0, 78}, {1, 10, 97}, {8, 14, 40}, {7, 9, 3}, {14, 6, 4}, {4, 8, 75}, {1, 6, 56}}, 1, 4, 10},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, findCheapestPrice(c.n, c.flights, c.src, c.dst, c.k))
		})
	}

}
