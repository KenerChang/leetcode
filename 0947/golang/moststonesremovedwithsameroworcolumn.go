package moststonesremovedwithsameroworcolumn

func find(un []int, i int) int {
	if un[i] == i {
		return i
	}

	un[i] = find(un, un[i])
	return un[i]
}

func isConntected(stoneA, stoneB []int) bool {
	return stoneA[0] == stoneB[0] || stoneA[1] == stoneB[1]
}

func removeStones(stones [][]int) int {
	// This is a undirected graph problem, the key to solve this problem is to find out
	// how many components in this graph
	// the answer is len(stones) - componentsCount, and we can get componentsCount
	// by using union find

	var componentsCount int = len(stones)

	// build graph
	un := make([]int, len(stones))
	sizes := make([]int, len(stones))
	for i := range stones {
		un[i] = i
		sizes[i] = 1
	}

	for i := 0; i < len(stones); i++ {
		for j := i + 1; j < len(stones); j++ {
			if isConntected(stones[i], stones[j]) {
				componentsCount -= union(un, sizes, i, j)
			}
		}
	}

	return len(stones) - componentsCount
}

func union(un, sizes []int, i, j int) int {
	pi := find(un, i)
	pj := find(un, j)

	if pi != pj {
		if sizes[pi] > sizes[pj] {
			un[pj] = pi
			sizes[pi] += sizes[pj]
		} else {
			un[pi] = pj
			sizes[pj] += sizes[pi]
		}
		return 1
	}
	return 0
}
