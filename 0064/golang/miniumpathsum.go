package miniumpathsum

import "fmt"

func minPathSum(grid [][]int) int {
	rowsNum := len(grid)
	if rowsNum == 0 {
		return 0
	}

	colsNum := len(grid[0])
	if colsNum == 0 {
		return 0
	}

	if rowsNum == 1 && colsNum == 1 {
		return grid[0][0]
	}

	pathCache := map[string]int{}
	return minPathSumImpl(grid, 0, 0, rowsNum-1, colsNum-1, pathCache)
}

func minPathSumImpl(grid [][]int, x, y, targetX, targetY int, pathCache map[string]int) int {
	// since we can only go right or go down
	// the minium path must through the right point or down point
	// so we can solve this problem by finding minium path of right point and down point recursively
	// until we reach the bottom-right point
	if x == targetX && y == targetY {
		return grid[x][y]
	}

	key := fmt.Sprintf("%d,%d", x, y)
	if minPath, ok := pathCache[key]; ok {
		fmt.Printf("hit cache, x: %d, y: %d", x, y)
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
