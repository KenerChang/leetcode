package numberofdistinctislands

func dfs(grids [][]int, r, c int) string {

	grids[r][c] = -1

	// search down
	var result string
	if r < len(grids)-1 && grids[r+1][c] == 1 {
		result += "D"
		result += dfs(grids, r+1, c)
		result += "-"
	}

	// search right
	if c < len(grids[0])-1 && grids[r][c+1] == 1 {
		result += "R"
		result += dfs(grids, r, c+1)
		result += "-"
	}

	// search up

	if r > 0 && grids[r-1][c] == 1 {
		result += "U"
		result += dfs(grids, r-1, c)
		result += "-"
	}

	// search left
	if c > 0 && grids[r][c-1] == 1 {
		result += "L"
		result += dfs(grids, r, c-1)
		result += "-"
	}

	return result
}

func numDistinctIslands(grids [][]int) int {
	// Use DFS

	hashMap := map[string]bool{}

	// iterate each grid
	var result int
	for r, row := range grids {
		for c, col := range row {
			if col == 1 {
				islandHash := dfs(grids, r, c)
				// fmt.Printf("r: %d, c: %d, islandHash: %s\n", r, c, islandHash)
				if _, ok := hashMap[islandHash]; !ok {
					result++
					hashMap[islandHash] = true
				}
			}
		}
	}
	return result
}
