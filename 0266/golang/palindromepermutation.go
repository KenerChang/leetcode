package palindromepermutation

func canPermutePalindrome(s string) bool {
	charBits := make([]int, 26)
	// for char in s
	for _, char := range s {
		// get index of char
		hashCode := hashChar(char)

		// flip chars[char]
		charBits[hashCode] = charBits[hashCode] ^ 1
	}

	// count number of not matched char
	result := 0
	for _, charBit := range charBits {
		result += charBit
	}

	// if s.length is an odd, or operation result should be 1
	// else s.length is na event, or operation result should be 0
	return result <= 1 && ((len(s) % 2) == result)
}

func hashChar(char rune) int {
	return int(char) - 97
}
