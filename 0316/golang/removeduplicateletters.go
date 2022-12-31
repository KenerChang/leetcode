package removeduplicateletters

func removeDuplicateLetters(s string) string {
	// count frequency of each character
	charsCount := make([]int, 26)
	for i := range s {
		c := s[i]
		charsCount[c-'a']++
	}

	stack := []byte{}

	used := map[byte]bool{}
	for i := range s {
		c := s[i]
		if used[c] {
			charsCount[c-'a']--
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > c && charsCount[stack[len(stack)-1]-'a'] > 0 {
			// pop chars which are larger than c and appear in later
			top := stack[len(stack)-1]
			delete(used, top)

			stack = stack[0 : len(stack)-1]
		}

		charsCount[c-'a']--
		stack = append(stack, c)
		used[c] = true
	}
	return string(stack)
}
