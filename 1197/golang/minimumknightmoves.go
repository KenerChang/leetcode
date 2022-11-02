package minimumknightmoves

func generateMovesFrom(x, y int) [][]int {

	return [][]int{
		{x - 1, y + 2}, // top left
		{x + 1, y + 2}, // top right
		{x - 2, y + 1}, // left top
		{x - 2, y - 1}, // left down
		{x - 1, y - 2}, // down left
		{x + 1, y - 2}, // down right
		{x + 2, y + 1}, // right top
		{x + 2, y - 1}, // right down
	}
}

func minKnightMoves(x int, y int) int {
	// We can find min moves by using BFS

	// init bitmap
	visited := make([][]bool, 700)
	for i := 0; i < 700; i++ {
		visited[i] = make([]bool, 700)
	}

	queue := [][]int{{0, 0, 0}} // 0: x, 1: y, 2: moves
	for len(queue) > 0 {
		// in each iteration:
		// check if current position is target position
		// generate moves from this position

		node := queue[0]
		queue = queue[1:]

		if node[0] == x && node[1] == y {
			return node[2]
		}

		moves := generateMovesFrom(node[0], node[1])
		for _, m := range moves {
			shiftX := m[0] + 350
			shiftY := m[1] + 350
			if visited[shiftX][shiftY] {
				continue
			}

			queue = append(queue, []int{m[0], m[1], node[2] + 1})
			visited[shiftX][shiftY] = true
		}
	}

	return -1
}
