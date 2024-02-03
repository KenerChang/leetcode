package removeinvalidparentheses

func removeInvalidParentheses(s string) []string {
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stack = append(stack, s[i])
		case ')':
			// pop stack if can match
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[0 : len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		}
	}

	if len(stack) == len(s) {
		return []string{""}
	}

	return backtracking(s, []byte{}, len(stack), 0, []string{}, map[string]bool{})
}

func backtracking(s string, now []byte, remainToRemove int, availableRightParentheses int, results []string, looked map[string]bool) []string {
	// reach end
	if s == "" {
		if availableRightParentheses == 0 {
			// a valid result
			result := string(now)
			if _, ok := looked[result]; !ok {
				results = append(results, result)
				looked[result] = true
			}
		}

		return results
	}

	charKept := append([]byte{}, now...)
	charKept = append(charKept, s[0])

	switch s[0] {
	case '(':
		// keep
		results = backtracking(s[1:], charKept, remainToRemove, availableRightParentheses+1, results, looked)

		// not keep
		if remainToRemove > 0 {
			results = backtracking(s[1:], now, remainToRemove-1, availableRightParentheses, results, looked)
		}
	case ')':
		// keep
		if availableRightParentheses > 0 {
			results = backtracking(s[1:], charKept, remainToRemove, availableRightParentheses-1, results, looked)
		}

		// not keep
		if remainToRemove > 0 {
			results = backtracking(s[1:], now, remainToRemove-1, availableRightParentheses, results, looked)
		}
	default:
		// keep all English letters
		results = backtracking(s[1:], charKept, remainToRemove, availableRightParentheses, results, looked)
	}

	return results
}
