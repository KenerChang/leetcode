package spiralmatrixII

func generateMatrix(n int) [][]int {
	total := n * n

	// init matrix
	output := make([][]int, n)
	for idx := range output {
		output[idx] = make([]int, n)
	}

	// define barrier
	lBarrier := -1 // left barrier
	rBarrier := n  // right barrier
	tBarrier := 0  // top barrier
	bBarrier := n  // bottom barrier
	direction := "right"
	changeDirection := false
	x, y := 0, 0

	count := 1
	for count <= total {
		if !changeDirection {
			output[x][y] = count
			count++
		}

		switch direction {
		case "left":
			if y-1 <= lBarrier {
				changeDirection = true
				bBarrier = x
				direction = "up"
			} else {
				changeDirection = false
				y--
			}
		case "right":
			if y+1 >= rBarrier {
				changeDirection = true
				tBarrier = x
				direction = "down"
			} else {
				y++
				changeDirection = false
			}
		case "up":
			if x-1 <= tBarrier {
				changeDirection = true
				lBarrier = y
				direction = "right"
			} else {
				changeDirection = false
				x--
			}
		case "down":
			if x+1 >= bBarrier {
				changeDirection = true
				rBarrier = y
				direction = "left"
			} else {
				changeDirection = false
				x++
			}
		}

	}
	return output
}
