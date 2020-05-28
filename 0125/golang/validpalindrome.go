package validpalindrome

// import (
// 	"fmt"
// )

func isPalindrome(s string) bool {
	fromHeadIdx := 0
	fromTailIdx := len(s) - 1

	var headChar, tailChar, diff int
	for fromHeadIdx < fromTailIdx {
		headChar = int(s[fromHeadIdx])
		tailChar = int(s[fromTailIdx])

		if !isAlphanumeric(headChar) {
			fromHeadIdx++
			continue
		}

		if !isAlphanumeric(tailChar) {
			fromTailIdx--
			continue
		}

		diff = headChar - tailChar
		if diff < 0 {
			diff = -diff
		}

		// fmt.Printf("headChar: %d, tailChar: %d, diff: %d", headChar, tailChar, diff)

		if diff != 0 && diff != 32 {
			return false
		}

		if diff == 32 && (headChar < 65 || tailChar < 65) {
			return false
		}

		fromHeadIdx++
		fromTailIdx--
	}
	return true
}

func isAlphanumeric(charCode int) bool {
	if charCode >= 48 && charCode <= 57 {
		// 0 ~ 9
		return true
	}

	if charCode >= 65 && charCode <= 90 {
		// A~Z
		return true
	}

	if charCode >= 97 && charCode <= 122 {
		return true
	}

	return false
}
