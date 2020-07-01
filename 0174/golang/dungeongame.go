package dungeongame

import (
	// "fmt"
	"math"
	"strconv"
)

func calculateMinimumHP(dungeon [][]int) int {
	cache := map[string]int{}
	neededHP := calculateMinimumHPImpl(dungeon, 0, 0, cache)

	if neededHP <= 0 {
		neededHP = -neededHP + 1
	} else {
		neededHP = 1
	}
	return neededHP
}

func calculateMinimumHPImpl(dungeon [][]int, rowIdx, colIdx int, cache map[string]int) int {
	endRow, endCol := len(dungeon)-1, len(dungeon[0])-1
	if rowIdx == endRow && colIdx == endCol {
		return min(0, dungeon[rowIdx][colIdx])
	}

	key := strconv.Itoa(rowIdx) + "," + strconv.Itoa(colIdx)
	if minNum, found := cache[key]; found {
		return minNum
	}

	var minRight, minDown int = math.MinInt32, math.MinInt32
	if rowIdx < endRow {
		minDown = calculateMinimumHPImpl(dungeon, rowIdx+1, colIdx, cache)
	}

	if colIdx < endCol {
		minRight = calculateMinimumHPImpl(dungeon, rowIdx, colIdx+1, cache)
	}

	minNum := max(minDown, minRight)
	minNum = min(0, dungeon[rowIdx][colIdx]+minNum)
	cache[key] = minNum

	return minNum
}

func max(num1, num2 int) int {
	if num1 >= num2 {
		return num1
	}
	return num2
}

func min(num1, num2 int) int {
	if num1 <= num2 {
		return num1
	}
	return num2
}
