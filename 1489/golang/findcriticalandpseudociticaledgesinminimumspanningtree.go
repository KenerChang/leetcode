package findcriticalandpseudociticaledgesinminimumspanningtree

import (
	"fmt"
	"sort"
)

type edges [][]int

func (e edges) Len() int {
	return len(e)
}

func (e edges) Less(i, j int) bool {
	return e[i][2] < e[j][2]
}

func (e edges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func findParent(parents []int, i int) int {
	if parents[i] != i {
		parents[i] = findParent(parents, parents[i])
	}

	return parents[i]
}

func union(parents, ranks []int, i, j int) {
	pi := findParent(parents, i)
	pj := findParent(parents, j)

	if pi == pj {
		return
	}

	if ranks[pi] > ranks[pj] {
		parents[pj] = pi
	} else if ranks[pi] < ranks[pj] {
		parents[pi] = pj
	} else {
		parents[pj] = pi
		ranks[pi]++
	}
}

func kruskals(n int, es [][]int, targetEdge int) (int, int) {
	// generate MST
	notUseWeight := 0

	// init union find
	parents := make([]int, n)
	ranks := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
		ranks[i] = 1
	}

	for i, e := range es {
		if i == targetEdge {
			continue
		}

		if findParent(parents, e[0]) != findParent(parents, e[1]) {
			notUseWeight += e[2]
			union(parents, ranks, e[0], e[1])
		}
	}

	// force to use target edge
	useWeight := 0
	if targetEdge >= 0 {
		// reset union find
		for i := 0; i < n; i++ {
			parents[i] = i
			ranks[i] = 1
		}

		useWeight += es[targetEdge][2]
		union(parents, ranks, es[targetEdge][0], es[targetEdge][1])

		for _, e := range es {
			if findParent(parents, e[0]) != findParent(parents, e[1]) {
				useWeight += e[2]
				union(parents, ranks, e[0], e[1])
			}
		}
	}

	return notUseWeight, useWeight
}

func findCriticalAndPseudoCriticalEdges(n int, es [][]int) [][]int {
	// Try to solve this problem by using Kruskal's algorithm
	// we can find a minimum spanning tree at first
	// then remove one edge a time from edges to see if the MST is changes:
	// 1. if we can not generate MST with the remaining edges, then it is a critical edge
	// 2. if we can generate another MST without the removed edge, then is is a pseudo critical edge

	// keep original index
	indexMap := map[string]int{}
	for i, e := range es {
		key := fmt.Sprintf("%d-%d", e[0], e[1])
		indexMap[key] = i
	}

	// sort edges
	sort.Sort(edges(es))

	// generate MST
	mstWeight, _ := kruskals(n, es, -1)

	// remove one edge a time
	criticals := []int{}
	pseudoCriticals := []int{}
	for i := 0; i < len(es); i++ {
		// generate MST without the removed edge

		notusedWeight, usedWeight := kruskals(n, es, i)

		key := fmt.Sprintf("%d-%d", es[i][0], es[i][1])
		if notusedWeight != mstWeight {
			// a critical edge
			criticals = append(criticals, indexMap[key])
		} else {
			if usedWeight == mstWeight {
				// a pseudo critical edge
				pseudoCriticals = append(pseudoCriticals, indexMap[key])
			}
		}
	}

	return [][]int{criticals, pseudoCriticals}
}
