package palindromepermutationii

import (
	"fmt"
	"sort"
)

var (
	charsMap []string = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
)

type sortRuneString []rune

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}

func generatePalindromes(s string) []string {
	// if not a palindromes, return empty string slice
	if !isPalindrome(s) {
		return []string{}
	}

	// sort string by characters
	runeArray := []rune(s)
	sort.Sort(sortRuneString(runeArray))

	// generate palindromes by sorted chars
	return generatePalindromesRecursive(string(runeArray), map[string][]string{})
}

func generatePalindromesRecursive(s string, resultsCache map[string][]string) []string {
	if len(s) == 0 {
		return []string{""}
	}
	// if s in cache, return results in cache
	if results, ok := resultsCache[s]; ok {
		return results
	}

	// iterate s from start
	i := 0
	results := []string{}
	inResults := map[string]bool{}
	for i < len(s) {
		// if i == len(s)-1 or s[i] != s[i+1], it is a single char, add char[i] to middle of results of s[0:i]+s[i+1:]
		if i == len(s)-1 || s[i] != s[i+1] {
			subString := getSubString(s, i, 1)
			subStringResults := generatePalindromesRecursive(subString, resultsCache)

			for _, subStringResult := range subStringResults {
				result := addCharToCenter(subStringResult, getChar(s[i]))
				if _, ok := inResults[result]; ok {
					continue
				}
				inResults[result] = true
				results = append(results, result)
			}
			i++
		} else {
			// else s[i] == s[i+1], they are a pair, add s[i] as prefix and suffix to results of s[0:i]+s[i+2:]
			subString := getSubString(s, i, 2)

			subStringResults := generatePalindromesRecursive(subString, resultsCache)
			for _, subStringResult := range subStringResults {
				result := fmt.Sprintf("%s%s%s", getChar(s[i]), subStringResult, getChar(s[i]))

				if _, ok := inResults[result]; ok {
					continue
				}
				inResults[result] = true

				results = append(results, result)
			}
			i += 2
		}
	}

	resultsCache[s] = results
	return results
}

func getSubString(s string, i int, qty int) string {
	if i == 0 {
		return s[i+qty:]
	} else {
		return s[0:i] + s[i+qty:]
	}
}

func getChar(char byte) string {
	pos := int(char) - 97
	return charsMap[int(pos)]
}

func addCharToCenter(s string, char string) string {
	middle := len(s) / 2
	return s[0:middle] + char + s[middle:]
}

func isPalindrome(s string) bool {
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
