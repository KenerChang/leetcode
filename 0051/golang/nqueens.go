package nqueens

func solveNQueens(n int) [][]string {
	if n == 1 {
		return [][]string{{"Q"}}
	}

	// init board
	board := make([][]int, n)
	for i := range board {
		board[i] = make([]int, n)
	}

	return nqueens(board, 0)
}

func nqueens(board [][]int, row int) [][]string {
	// if row equals len(board) -1,
	// return all available positions in this row
	if row == len(board)-1 {
		columns := board[row]

		results := [][]string{}
		for i := range columns {
			if canPutOnPosition(board, row, i) {
				// fmt.Printf("last row: %d, column: %d\n", row, i)
				result := generageRowCandiate(len(board), i)
				results = append(results, []string{result})
			}
		}

		return results
	}

	// check if row has position to put a queen
	results := [][]string{}
	for i := 0; i < len(board[row]); i++ {
		if !canPutOnPosition(board, row, i) {
			continue
		}

		// disable all positions which the queen can attack
		putOnPosition(board, row, i)
		// fmt.Printf("board: %+v, row: %d, column: %d\n", board, row, i)

		candidates := nqueens(board, row+1)

		for _, candidate := range candidates {
			rowResult := generageRowCandiate(len(board), i)
			result := append([]string{rowResult}, candidate...)
			results = append(results, result)
		}

		// reset board
		resetBoard(board, row, i)
	}

	return results
}

func canPutOnPosition(board [][]int, row, column int) bool {
	return board[row][column] == 0
}

func putOnPosition(board [][]int, row, column int) {
	board[row][column] += 1

	// disable the row
	for i := range board[row] {
		board[row][i] += 1
	}

	// fmt.Printf("board: %+v\n", board)

	// disable the column
	for i := range board {
		board[i][column] += 1
	}

	// disable diagonals
	// left top
	i, j := row-1, column-1
	for i >= 0 && j >= 0 {
		board[i][j] += 1
		i--
		j--
	}

	// right top
	i, j = row-1, column+1
	for i >= 0 && j < len(board) {
		board[i][j] += 1
		i--
		j++
	}

	// left bottom
	i, j = row+1, column-1
	for i < len(board) && j >= 0 {
		board[i][j] += 1
		i++
		j--
	}

	// right bottom
	i, j = row+1, column+1
	for i < len(board) && j < len(board) {
		board[i][j] += 1
		i++
		j++
	}
}

func generageRowCandiate(rowLength, column int) string {
	result := make([]byte, rowLength)
	for i := range result {
		result[i] = '.'
	}
	result[column] = 'Q'
	return string(result)
}

func resetBoard(board [][]int, row, column int) {
	board[row][column] -= 1

	// reset rows
	for i := range board[row] {
		board[row][i] -= 1
	}

	// reset columns
	for i := range board {
		board[i][column] -= 1
	}

	// reset diagonals
	// left top
	i, j := row-1, column-1
	for i >= 0 && j >= 0 {
		board[i][j] -= 1
		i--
		j--
	}

	// right top
	i, j = row-1, column+1
	for i >= 0 && j < len(board) {
		board[i][j] -= 1
		i--
		j++
	}

	// left bottom
	i, j = row+1, column-1
	for i < len(board) && j >= 0 {
		board[i][j] -= 1
		i++
		j--
	}

	// right bottom
	i, j = row+1, column+1
	for i < len(board) && j < len(board) {
		board[i][j] -= 1
		i++
		j++
	}
}
