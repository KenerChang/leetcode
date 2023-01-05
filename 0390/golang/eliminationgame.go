package eliminationgame

func lastRemaining(n int) int {
	remaining := n
	head := 1
	step := 1
	left := true

	for remaining > 1 {
		if left || remaining%2 == 1 {
			head += step
		}

		remaining = remaining / 2
		step *= 2
		left = !left
	}

	return head
}
