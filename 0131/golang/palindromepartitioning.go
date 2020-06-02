package palindromepartitioning

import (
	"strings"
)

func partition(s string) [][]string {
	result := [][]string{
		make([]string, len(s)),
	}

	if len(s) == 1 {
		return [][]string{
			[]string{s},
		}
	}

	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			result[0][i] = s[i:]
		} else {
			result[0][i] = s[i : i+1]
		}
	}

	key := strings.Join(result[0], "|")

	cache := map[string]bool{
		key: true,
	}
	palindromes := partitionImpl(s, cache)
	result = append(result, palindromes...)

	return result
}

func partitionImpl(s string, cache map[string]bool) [][]string {
	if len(s) == 1 {
		return [][]string{
			[]string{s},
		}
	}

	result := [][]string{}
	for i := 1; i < len(s)+1; i++ {
		if i == len(s) {
			isp := isPalindrome(s)
			if isp {
				result = append(result, []string{s})
			}
			cache[s] = isp
		}

		firstPart := s[0:i]
		secondPart := s[i:]

		if isp, found := cache[firstPart]; found {
			if !isp {
				continue
			}
		}

		isp := isPalindrome(firstPart)
		cache[firstPart] = isp

		if !isp {
			continue
		}

		secondPartPalindromes := partitionImpl(secondPart, cache)
		for _, palindrome := range secondPartPalindromes {
			withFirstPart := append([]string{firstPart}, palindrome...)
			key := strings.Join(withFirstPart, "|")

			if _, found := cache[key]; found {
				continue
			}

			result = append(result, withFirstPart)
		}
	}
	return result
}

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
