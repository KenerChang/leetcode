package miniumpathsum

// import "fmt"

func minPathSumDP(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	targetX := len(grid) - 1
	targetY := len(grid[0]) - 1

	minPath := grid[targetX][targetY]
	for i := targetX - 1; i >= 0; i-- {
		minPath += grid[i][targetY]
		grid[i][targetY] = minPath
	}

	minPath = grid[targetX][targetY]
	for i := targetY - 1; i >= 0; i-- {
		minPath += grid[targetX][i]
		grid[targetX][i] = minPath
	}

	for i := targetX - 1; i >= 0; i-- {
		for j := targetY - 1; j >= 0; j-- {
			min := Min(grid[i+1][j], grid[i][j+1])
			grid[i][j] = min + grid[i][j]
		}
	}
	return grid[0][0]
}

func minPathSum(grid [][]int) int {
	rowsNum := len(grid)
	if rowsNum == 0 {
		return 0
	}
	targetX := rowsNum - 1

	colsNum := len(grid[0])
	if colsNum == 0 {
		return 0
	}
	targetY := colsNum - 1

	if rowsNum == 1 && colsNum == 1 {
		return grid[0][0]
	}

	pathCache := map[int64]int{}
	// compute some default cache
	minPath := grid[targetX][targetY]
	for i := targetX - 1; i >= 0; i-- {
		minPath += grid[i][targetY]
		key := getKey(i, targetY)
		pathCache[key] = minPath
	}

	minPath = grid[targetX][targetY]
	for i := targetY - 1; i >= 0; i-- {
		minPath += grid[targetX][i]
		key := getKey(targetX, i)
		pathCache[key] = minPath
	}

	return minPathSumImpl(grid, 0, 0, targetX, targetY, pathCache)
}

func minPathSumImpl(grid [][]int, x, y, targetX, targetY int, pathCache map[int64]int) int {
	// since we can only go right or go down
	// the minium path must through the right point or down point
	// so we can solve this problem by finding minium path of right point and down point recursively
	// until we reach the bottom-right point
	if x == targetX && y == targetY {
		return grid[x][y]
	}

	key := getKey(x, y)
	if minPath, ok := pathCache[key]; ok {
		return minPath
	}

	minPaths := []int{}
	if x+1 < len(grid) {
		// min path of choosing to go down
		minPaths = append(minPaths, minPathSumImpl(grid, x+1, y, targetX, targetY, pathCache))
	}

	if y+1 < len(grid[0]) {
		// min path of choosing to go right
		minPaths = append(minPaths, minPathSumImpl(grid, x, y+1, targetX, targetY, pathCache))
	}

	minPath := Min(minPaths...)
	minPath += grid[x][y]

	pathCache[key] = minPath
	return minPath
}

func Min(nums ...int) int {
	if nums == nil {
		return 0
	}

	min := 0
	for idx, num := range nums {
		if idx == 0 || num < min {
			min = num
		}
	}
	return min
}

func getKey(x, y int) int64 {
	key := int64(x)
	key = key << 32
	key += int64(y)
	return key
}
