package stringcompression

import "strconv"

func compress(chars []byte) int {
	// Since chars is a consecutive repeating characters, we can use a pointer
	// to keep its position

	var charPos int
	charCount := 1
	var curChar byte = chars[0]

	for i := 1; i < len(chars); i++ {
		if chars[i] != curChar {
			// char changed, need to place to result
			charPos = updateChars(chars, curChar, charPos, charCount)

			charCount = 1
			curChar = chars[i]
		} else {
			charCount++
		}
	}

	charPos = updateChars(chars, curChar, charPos, charCount)
	return charPos
}

func updateChars(chars []byte, curChar byte, charPos, charCount int) int {
	chars[charPos] = curChar

	if charCount == 1 {
		charPos++
	} else {
		charPos++

		charCountStr := strconv.Itoa(charCount)
		for _, r := range charCountStr {
			chars[charPos] = byte(r)
			charPos++
		}
	}

	return charPos
}
