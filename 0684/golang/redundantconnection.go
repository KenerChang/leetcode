package redundantconnection

func findRedundantConnection(edges [][]int) []int {
	// Use DSU (disjoint set union)

	roots := make([]int, len(edges)+1)
	ranks := make([]int, len(edges)+1)
	for i := range roots {
		roots[i] = i
		ranks[i] = 1
	}

	for _, edge := range edges {
		if !union(roots, ranks, edge[0], edge[1]) {
			return edge
		}
	}

	return []int{}
}

func find(roots []int, i int) int {
	if roots[i] == i {
		return i
	}

	roots[i] = find(roots, roots[i])
	return roots[i]
}

func union(roots, ranks []int, i, j int) bool {
	rooti := find(roots, i)
	rootj := find(roots, j)

	if rooti == rootj {
		return false
	}

	if ranks[rooti] <= ranks[rootj] {
		roots[rooti] = rootj
		ranks[rootj] += ranks[rooti]
	} else {
		roots[rootj] = rooti
		ranks[rooti] += ranks[rootj]
	}

	return true
}
