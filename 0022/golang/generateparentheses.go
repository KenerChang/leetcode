package generateparentheses

func generateParenthesis(n int) []string {
	// Use backtracking to solve this problem

	return backtracking(n, 0, 0, 0, make([]byte, 2*n), []string{}, map[string]struct{}{})
}

func backtracking(n, leftP, rightP, currentPos int, current []byte, answers []string, seen map[string]struct{}) []string {
	// check if reach the end
	if currentPos == 2*n {
		answer := string(current)
		if _, ok := seen[answer]; ok {
			return answers
		}

		answers = append(answers, string(answer))
		seen[answer] = struct{}{}
		return answers
	}

	// check if we can put left parenthesis
	if leftP < n {
		current[currentPos] = '('
		answers = backtracking(n, leftP+1, rightP, currentPos+1, current, answers, seen)
	}

	// check if we can put right parenthesis
	if leftP > rightP {
		current[currentPos] = ')'
		answers = backtracking(n, leftP, rightP+1, currentPos+1, current, answers, seen)
	}

	return answers
}
