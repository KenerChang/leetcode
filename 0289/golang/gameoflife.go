package gameoflife

func getState(board [][]int, row, col int) int {
	var neighborLiveCount int

	// get up
	if row > 0 {
		neighborLiveCount += board[row-1][col] % 2
	}

	// get down
	if row < len(board)-1 {
		neighborLiveCount += board[row+1][col] % 2
	}

	// get left
	if col > 0 {
		neighborLiveCount += board[row][col-1] % 2
	}

	// get right
	if col < len(board[0])-1 {
		neighborLiveCount += board[row][col+1] % 2
	}

	// get up left
	if row > 0 && col > 0 {
		neighborLiveCount += board[row-1][col-1] % 2
	}

	// get up right
	if row > 0 && col < len(board[0])-1 {
		neighborLiveCount += board[row-1][col+1] % 2
	}

	// get down left
	if row < len(board)-1 && col > 0 {
		neighborLiveCount += board[row+1][col-1] % 2
	}

	// get down right
	if row < len(board)-1 && col < len(board[0])-1 {
		neighborLiveCount += board[row+1][col+1] % 2
	}

	// check rules
	if board[row][col]%2 == 0 {
		// now is dead
		// check rule 4

		if neighborLiveCount == 3 {
			return 2
		}
		return 0
	} else {
		// alive
		if neighborLiveCount < 2 {
			return 3
		} else if neighborLiveCount > 3 {
			return 3
		} else {
			return 1
		}
	}
}

func gameOfLife(board [][]int) {
	// We can solve this problem in 2 steps

	// 1. mark state
	// 0: now is 0, next generation is 0
	// 1: now is 1, next generation is 1
	// 2: now is 0, next generation is 1
	// 3: now is 1, next generation is 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			// fmt.Printf("i: %d, j: %d, board[%d][%d]: %d, state: %d\n", i, j, i, j, board[i][j], getState(board, i, j))
			board[i][j] = getState(board, i, j)
		}
	}

	// 2. update state
	// change marked state to alive or dead
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 2 {
				board[i][j] = 1
			} else if board[i][j] == 3 {
				board[i][j] = 0
			}
		}
	}
}
