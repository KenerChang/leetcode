package dungeongame

// import (
// 	// "fmt"
// 	"math"
// 	// "strconv"
// )

func calculateMinimumHP(dungeon [][]int) int {
	endRow, endCol := len(dungeon)-1, len(dungeon[0])-1

	var minNum int
	for i := endRow; i >= 0; i-- {
		for j := endCol; j >= 0; j-- {
			if i == endRow && j == endCol {
				minNum = min(0, dungeon[i][j])
			} else if i == endRow {
				minNum = min(0, dungeon[i][j+1]+dungeon[i][j])
			} else if j == endCol {
				minNum = min(0, dungeon[i+1][j]+dungeon[i][j])
			} else {
				minNum = max(dungeon[i+1][j], dungeon[i][j+1])
				minNum = dungeon[i][j] + minNum
				minNum = min(0, minNum)
			}

			dungeon[i][j] = minNum
		}
	}

	minNum = dungeon[0][0]
	if minNum <= 0 {
		minNum = -minNum + 1
	} else {
		minNum = 1
	}

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
