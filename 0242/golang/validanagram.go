package validanagram

func isAnagram(s string, t string) bool {
	charMap := map[rune]int{}

	for _, char := range s {
		charMap[char]++
	}

	for _, char := range t {
		count, ok := charMap[char]
		if !ok {
			return false
		}

		count--
		charMap[char] = count
		if count == 0 {
			delete(charMap, char)
		}
	}

	return len(charMap) == 0
}
