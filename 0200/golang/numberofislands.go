package numberofislands

func numIslands(grid [][]byte) int {
	count := 0
	for rowIdx, row := range grid {
		for colIdx, char := range row {
			if char == '1' {
				count++

				setSearched(grid, rowIdx, colIdx)
			}
		}
	}
	return count
}

func setSearched(grid [][]byte, row, col int) {
	rowSize := len(grid)
	colSize := len(grid[0])

	if row >= rowSize || row < 0 || col < 0 || col >= colSize {
		return
	}

	if grid[row][col] != '1' {
		return
	}

	grid[row][col] = 'X'

	setSearched(grid, row-1, col)
	setSearched(grid, row+1, col)
	setSearched(grid, row, col-1)
	setSearched(grid, row, col+1)
}
