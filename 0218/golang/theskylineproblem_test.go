package theskylineproblem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSkyline(t *testing.T) {
	cases := []struct {
		name     string
		building [][]int
		want     [][]int
	}{
		{"[[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]", [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}, [][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}}},
		{"[[0,2,3],[2,5,3]]", [][]int{{0, 2, 3}, {2, 5, 3}}, [][]int{{0, 3}, {5, 0}}},
		{"[[1,2,1],[1,2,2],[1,2,3]]", [][]int{{1, 2, 1}, {1, 2, 2}, {1, 2, 3}}, [][]int{{1, 3}, {2, 0}}},
		{"[[0,3,3],[1,5,3],[2,4,3],[3,7,3]]", [][]int{{0, 3, 3}, {1, 5, 3}, {2, 4, 3}, {3, 7, 3}}, [][]int{{0, 3}, {7, 0}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.want, getSkyline(c.building))
		})
	}
}
