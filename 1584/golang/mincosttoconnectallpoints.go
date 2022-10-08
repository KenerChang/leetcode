package mincosttoconnectallpoints

import (
	"math"
	"sort"
)

type edge struct {
	u int
	v int
	w float64
}

type edges []edge

func (e edges) Len() int {
	return len(e)
}

func (e edges) Less(i, j int) bool {
	return e[i].w < e[j].w
}

func (e edges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func minCostConnectPoints(points [][]int) int {
	// This problem is a minimum spanning tree problem.
	// We can use Kruskal's algorithm to solve it.

	// collect edges
	es := []edge{}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			es = append(es, edge{
				u: i,
				v: j,
				w: math.Abs(float64(points[i][0]-points[j][0])) + math.Abs(float64(points[i][1]-points[j][1])),
			})
		}
	}

	// sort edges
	sort.Sort(edges(es))

	// build parents and ranks
	parents := make([]int, len(points))
	ranks := make([]int, len(points))
	for i := 0; i < len(points); i++ {
		parents[i] = i
		ranks[i] = 0
	}

	// iterate edges
	var result float64
	for _, e := range es {
		if parent(parents, e.u) != parent(parents, e.v) {
			union(parents, ranks, e.u, e.v)

			result += e.w
		}
	}

	return int(result)
}

func parent(parents []int, n int) int {
	if n == parents[n] {
		return n
	}

	p := parent(parents, parents[n])
	parents[n] = p

	return p
}

func union(parents, ranks []int, u, v int) {
	pu := parent(parents, u)
	pv := parent(parents, v)

	if pu != pv {
		// do merge
		if ranks[pu] < ranks[pv] {
			parents[pv] = parents[pu]
		} else {
			parents[pu] = parents[pv]
		}

		if ranks[pu] == ranks[pv] {
			ranks[pv]++
		}
	}
}
