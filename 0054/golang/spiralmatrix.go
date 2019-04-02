package spiralmatrix

import "fmt"

func spiralOrder(matrix [][]int) []int {
	// 1. define barrier
	// 2. when hit barrier change direction

	// empty matrix
	if len(matrix) == 0 {
		return []int{}
	}

	// no elements in rows
	if len(matrix[0]) == 0 {
		return []int{}
	}

	tBarrier := -1             // top barrier
	bBarrier := len(matrix)    // bottom barrier
	lBarrier := -1             // left barrier
	rBarrier := len(matrix[0]) // right barrier

	output := []int{}
	total := len(matrix) * len(matrix[0])

	xIndex, yIndex := 0, 0
	direction := "right"
	changeDirection := false
	for len(output) < total {
		if !changeDirection {
			fmt.Printf("x: %d, y: %d\n", xIndex, yIndex)
			output = append(output, matrix[xIndex][yIndex])
		}

		switch direction {
		case "left":
			if yIndex-1 <= lBarrier {
				direction = "up"
				bBarrier = xIndex
				changeDirection = true
			} else {
				yIndex--
				changeDirection = false
			}
		case "right":
			if yIndex+1 >= rBarrier {
				direction = "down"
				tBarrier = xIndex
				changeDirection = true
			} else {
				yIndex++
				changeDirection = false
			}
		case "up":
			if xIndex-1 <= tBarrier {
				direction = "right"
				lBarrier = yIndex
				changeDirection = true
			} else {
				xIndex--
				changeDirection = false
			}
		case "down":
			if xIndex+1 >= bBarrier {
				direction = "left"
				rBarrier = yIndex
				changeDirection = true
			} else {
				xIndex++
				changeDirection = false
			}
		}
	}
	return output
}
