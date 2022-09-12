package flipgameii

func canWin(currentState string) bool {
	// find a move that can guarantee a win
	// it means that we are finding if there exists a move that
	// the opponent cannot win

	childrenMoves := getChildren(currentState)
	if len(childrenMoves) == 0 {
		return false // player1 cannot do any move
	}

	cache := make(map[string]bool)
	for _, child := range childrenMoves {
		if !recursive(child, cache) {
			// the opponent can win
			return true
		}
	}

	return false
}

func recursive(state string, cache map[string]bool) bool {
	// this util function returns if there exists any possibility that that player can win
	if cw, ok := cache[state]; ok {
		return cw
	}

	childrenMoves := getChildren(state)
	if len(childrenMoves) == 0 {
		cache[state] = false
		return false
	}

	var anywin bool
	for _, child := range childrenMoves {
		win := recursive(child, cache)

		anywin = anywin || !win // anywin is true if at least one child can win
	}

	cache[state] = anywin
	return anywin
}

func getChildren(state string) []string {
	results := []string{}
	for i := 0; i < len(state)-1; i++ {
		if state[i] == '+' && state[i+1] == '+' {
			var childStateString []byte
			if i == 0 {
				childStateString = append([]byte{'-', '-'}, state[i+2:]...)
			} else if i == len(state)-2 {
				childStateString = append([]byte(state[:i]), '-', '-')
			} else {
				// childStateString = append([]byte(state[:i]), '-', '-', state[i+2:]...)
				childStateString = []byte(state[:i])
				suffix := []byte{'-', '-'}
				childStateString = append(childStateString, suffix...)
				childStateString = append(childStateString, state[i+2:]...)
			}

			results = append(results, string(childStateString))
		}
	}
	return results
}
