package uniquepathsii

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}

	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}

	// check if start point is available
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	// check if target is available
	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	cache := map[string]int{}
	return uniquePathsWithObstaclesImpl(obstacleGrid, 0, 0, m-1, n-1, cache)
}

func uniquePathsWithObstaclesImpl(obstacleGrid [][]int, x, y, targetX, targetY int, pathsCache map[string]int) int {
	if x == targetX && y == targetY {
		if obstacleGrid[x][y] == 0 {
			return 1
		} else {
			return 0
		}
	}

	if (x+1 == targetX && y == targetY) || (x == targetX && y+1 == targetY) {
		return 1
	}

	key := fmt.Sprintf("%d,%d", x, y)
	if paths, ok := pathsCache[key]; ok {
		return paths
	}

	var paths int
	if x < targetX && obstacleGrid[x+1][y] != 1 {
		paths += uniquePathsWithObstaclesImpl(obstacleGrid, x+1, y, targetX, targetY, pathsCache)
	}

	if y < targetY && obstacleGrid[x][y+1] != 1 {
		paths += uniquePathsWithObstaclesImpl(obstacleGrid, x, y+1, targetX, targetY, pathsCache)
	}

	// set cache
	pathsCache[key] = paths

	return paths
}
