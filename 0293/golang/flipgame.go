package flipgame

func generatePossibleNextMoves(currentState string) []string {
	bytes := []byte(currentState)

	results := []string{}
	for i := 0; i < len(bytes)-1; i++ {
		if bytes[i] == '+' && bytes[i+1] == '+' {
			bytes[i] = '-'
			bytes[i+1] = '-'

			results = append(results, string(bytes))

			bytes[i] = '+'
			bytes[i+1] = '+'
		}
	}

	return results
}
