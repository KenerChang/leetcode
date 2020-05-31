package surroundedregions

import (
	// "fmt"
	"strconv"
)

var X byte = byte('X')
var O byte = byte('O')

func solve(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}

	xdimension := len(board)
	ydimension := len(board[0])

	for x, row := range board {
		for y, cell := range row {
			if cell == X {
				continue
			}

			if !isBoarder(x, y, xdimension, ydimension) {
				continue
			}

			if cell != O {
				continue
			}

			// handle O which has not handled
			flipCandidates := getFlipCandidates(board, x, y, xdimension, ydimension)

			for _, candidate := range flipCandidates {
				xposition, yposition := candidate[0], candidate[1]
				board[xposition][yposition] = '#'
			}
		}
	}

	for x, row := range board {
		for y, cell := range row {
			if cell == '#' {
				board[x][y] = O
			} else if cell == O {
				board[x][y] = X
			}
		}
	}
}

func getFlipCandidates(board [][]byte, x, y, xdimension, ydimension int) [][]int {
	queue := [][]int{
		[]int{x - 1, y},
		[]int{x + 1, y},
		[]int{x, y - 1},
		[]int{x, y + 1},
	}

	candidates := [][]int{
		[]int{x, y},
	}
	visited := map[string]bool{}
	for len(queue) > 0 {
		position := queue[0]
		queue = queue[1:]
		xposition, yposition := position[0], position[1]

		if xposition < 0 || yposition < 0 || xposition >= xdimension || yposition >= ydimension {
			continue
		}

		if xposition == x && yposition == y {
			continue
		}

		key := strconv.Itoa(xposition) + "," + strconv.Itoa(yposition)
		if _, found := visited[key]; found {
			continue
		}

		if board[xposition][yposition] == O {
			queue = append(queue,
				[]int{xposition - 1, yposition},
				[]int{xposition + 1, yposition},
				[]int{xposition, yposition - 1},
				[]int{xposition, yposition + 1},
			)
			candidates = append(candidates, []int{xposition, yposition})
			visited[key] = true
		}
	}
	return candidates
}

func isBoarder(x, y, xdimension, ydimension int) bool {
	if x == 0 || y == 0 || x == xdimension-1 || y == ydimension-1 {
		return true
	}
	return false
}
