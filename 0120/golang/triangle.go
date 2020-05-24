package triangle

// import (
// 	"math"
// )

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	levelCache := map[int64]int{}

	return minimumTotalImpl(triangle, 0, 0, levelCache)
}

func minimumTotalImpl(triangle [][]int, xshift, yshift int, levelCache map[int64]int) int {
	if yshift >= len(triangle) {
		return 0
	}

	key := int64(xshift)<<32 + int64(yshift)
	if minPath, found := levelCache[key]; found {
		return minPath
	}

	leftMinPath := minimumTotalImpl(triangle, xshift, yshift+1, levelCache)
	rightMinPath := minimumTotalImpl(triangle, xshift+1, yshift+1, levelCache)

	minPath := triangle[yshift][xshift] + min(leftMinPath, rightMinPath)

	levelCache[key] = minPath

	return minPath
}

func min(num1, num2 int) int {
	if num1 <= num2 {
		return num1
	}
	return num2
}
