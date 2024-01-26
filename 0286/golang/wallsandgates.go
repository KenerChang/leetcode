package wallsandgates

const (
	FROM_UPPER = iota
	FROM_BOTTOM
	FROM_LEFT
	FROM_RIGHT
)

func wallsAndGates(rooms [][]int) {
	rowCount := len(rooms)
	colCount := len(rooms[0])

	// iterate all cell
	for m := 0; m < rowCount; m++ {
		for n := 0; n < colCount; n++ {
			switch rooms[m][n] {
			case 0:
				// a gate, update shortest pathes
				updateShortestPath(rooms, m, n)
			default:
				// do nothing
			}
		}
	}
}

func updateShortestPath(rooms [][]int, m, n int) {
	// start from rooms[m][n] and do dfs to find shortest path
	// from the room

	// go up
	dfs(rooms, m, n-1, FROM_BOTTOM, 1)

	// go down
	dfs(rooms, m, n+1, FROM_UPPER, 1)

	// go left
	dfs(rooms, m-1, n, FROM_RIGHT, 1)

	// go right
	dfs(rooms, m+1, n, FROM_LEFT, 1)
}

func dfs(rooms [][]int, m, n, from, steps int) {
	rowCount := len(rooms)
	colCount := len(rooms[0])

	if m >= rowCount || n >= colCount || m < 0 || n < 0 {
		// out of boundary
		return
	}

	if rooms[m][n] <= 0 {
		// gate or obstacle
		return
	}

	// a room, update steps if is a shorter path
	if steps >= rooms[m][n] {
		return
	}

	rooms[m][n] = steps

	if from != FROM_UPPER {
		// go up
		dfs(rooms, m, n-1, FROM_BOTTOM, steps+1)
	}

	if from != FROM_BOTTOM {
		// go down
		dfs(rooms, m, n+1, FROM_UPPER, steps+1)
	}

	if from != FROM_LEFT {
		// go left
		dfs(rooms, m-1, n, FROM_RIGHT, steps+1)
	}

	if from != FROM_RIGHT {
		// go right
		dfs(rooms, m+1, n, FROM_LEFT, steps+1)
	}
}
