package rotatestring

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	// if goal[i] == s[0], check if goal is a rotation of s
	for i, c := range goal {
		if c == rune(s[0]) {
			if goal[i:]+goal[0:i] == s {
				return true
			}
		}
	}

	return false
}
