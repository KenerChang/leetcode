package removecoloredpiecesifbothneighborsarethesamecolor

func winnerOfGame(colors string) bool {
	// count move counts that A & B can take
	// so if A can take more moves, A wins
	// else B wins

	var amoves, bmoves, lastIdx int
	var last rune

	for i, char := range colors {
		if char != last {
			// character changes
			if i == 0 {
				last = char
			} else {
				// count move
				if last == 'A' {
					amoves += max(0, i-lastIdx-2)
				} else {
					bmoves += max(0, i-lastIdx-2)
				}
				last = char
				lastIdx = i
			}
		}
	}

	if lastIdx != len(colors)-1 {
		if last == 'A' {
			amoves += max(0, len(colors)-lastIdx-2)
		} else {
			bmoves += max(0, len(colors)-lastIdx-2)
		}
	}

	return amoves > bmoves
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
