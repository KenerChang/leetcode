package ransomnote

func canConstruct(ransomNote string, magazine string) bool {
	charCount := make([]int, 26)

	for _, c := range magazine {
		charCount[c-'a']++
	}

	for _, c := range ransomNote {
		idx := c - 'a'
		if charCount[idx] == 0 {
			return false
		}

		charCount[idx]--
	}

	return true
}
