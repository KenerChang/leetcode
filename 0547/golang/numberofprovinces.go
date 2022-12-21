package numberofprovinces

func findCircleNum(isConnected [][]int) int {
	// We can use union find to group cities
	// so cities which belong to same parent denote that they are in same province

	// parents
	parents := make([]int, len(isConnected))
	ranks := make([]int, len(isConnected))
	for i := range isConnected {
		parents[i] = i
		ranks[i] = 1
	}

	for i := range isConnected {
		for j := range isConnected {
			if isConnected[i][j] == 1 {
				union(parents, ranks, i, j)
			}
		}
	}

	provinceMap := map[int]bool{}
	for i := range parents {
		find(parents, i)
		provinceMap[parents[i]] = true
	}

	return len(provinceMap)
}

func find(parents []int, i int) int {
	if parents[i] == i {
		return i
	}

	parents[i] = find(parents, parents[i])
	return parents[i]
}

func union(parents, ranks []int, i, j int) {
	pari := find(parents, i)
	parj := find(parents, j)

	if pari != parj {
		if ranks[pari] <= ranks[parj] {
			parents[pari] = parj
			ranks[parj] += ranks[pari]
		} else {
			parents[parj] = pari
			ranks[pari] += ranks[parj]
		}
	}
}
