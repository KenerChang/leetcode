package maxareaofisland

func dfs(grids [][]int, r, c, area int) int {
	grids[r][c] = -1

	area++

	// visit left
	if c > 0 && grids[r][c-1] == 1 {
		area = dfs(grids, r, c-1, area)
	}

	// visit right
	if c < len(grids[0])-1 && grids[r][c+1] == 1 {
		area = dfs(grids, r, c+1, area)
	}

	// visit up
	if r > 0 && grids[r-1][c] == 1 {
		area = dfs(grids, r-1, c, area)
	}

	// visit down
	if r < len(grids)-1 && grids[r+1][c] == 1 {
		area = dfs(grids, r+1, c, area)
	}

	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxAreaOfIsland(grids [][]int) int {
	// Use DFS

	var result int
	for r, row := range grids {
		for c, col := range row {
			if col == 1 {
				result = max(result, dfs(grids, r, c, 0))
			}
		}
	}
	return result
}
