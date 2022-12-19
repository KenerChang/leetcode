package firstuniquecharacterinastring

func firstUniqChar(s string) int {
	charFrequency := make([]int, 26)

	for _, char := range s {
		charFrequency[char-'a']++
	}

	for i, char := range s {
		if charFrequency[char-'a'] == 1 {
			return i
		}
	}

	return -1
}
