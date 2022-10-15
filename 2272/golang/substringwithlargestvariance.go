package substringwithlargestvariance

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestVariance(s string) int {
	// We can solve this problem with Kadane's algorithm.
	// To reduce the problem into max subarray problem.
	// we compare any two characters in the string ang get the max variance between these 2 characters.

	// build character set
	charSet := make(map[rune]int)
	for _, c := range s {
		charSet[c]++
	}

	maxVar := 0
	for a := range charSet {
		for b := range charSet {
			if a == b {
				continue
			}

			// the problem is how we have a efficent solveOnce function
			variance := solveOnce(a, b, s)
			// fmt.Printf("a: %c, b: %c, variance: %d\n", a, b, variance)
			maxVar = max(maxVar, variance)
		}
	}
	return maxVar
}

func solveOnce(a, b rune, s string) int {
	// Use Kanade's algorithm, we can solve the problem in O(n)
	maxVar := 0
	variance := 0
	var hasB bool
	var firstB bool
	for i := 0; i < len(s); i++ {
		if rune(s[i]) == a {
			variance++
		} else if rune(s[i]) == b {
			hasB = true

			if firstB && variance >= 0 {
				// we have a ba+ pattern, so reset the firstB to indicate it is now like a+b pattern
				firstB = false
			} else if (variance - 1) < 0 {
				// not a ba+ pattern, and also not has any a, so we can start a ba+ pattern
				firstB = true
				variance = -1
			} else {
				// maybe a a+ pattern, then we add a b and minus 1 to variance
				variance -= 1
			}
		}

		if hasB {
			maxVar = max(maxVar, variance)
		}

	}
	return maxVar
}
