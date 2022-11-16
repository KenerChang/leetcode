package rottingoranges

func orangesRotting(grid [][]int) int {
	// We can solve this problem by iterating grid and mark oranges as rotten until converge\
	// For peventing mess up rotten oranges & going to be rotten oranges
	// we need an extra state to indicate an orange is going to be rotten

	var changed bool = true
	var result int = -1
	var notRottenOrangeCount int

	for changed {
		changed = false
		toBeRottenLists := [][]int{}
		notRottenOrangeCount = 0
		result++
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				if grid[i][j] == 0 || grid[i][j] == 2 {
					continue
				}

				willRotten := becomeRotten(grid, i, j)
				changed = changed || willRotten

				if willRotten {
					toBeRottenLists = append(toBeRottenLists, []int{i, j})
					grid[i][j] = 3
				} else {
					notRottenOrangeCount++
				}
			}
		}

		for _, pos := range toBeRottenLists {
			grid[pos[0]][pos[1]] = 2
		}
	}

	if notRottenOrangeCount > 0 {
		return -1
	}

	return result
}

func becomeRotten(grid [][]int, row, col int) bool {
	// up
	if row > 0 && grid[row-1][col] == 2 {
		return true
	}

	// bottom
	if row < len(grid)-1 && grid[row+1][col] == 2 {
		return true
	}

	// left
	if col > 0 && grid[row][col-1] == 2 {
		return true
	}

	// right
	if col < len(grid[0])-1 && grid[row][col+1] == 2 {
		return true
	}

	return false
}
