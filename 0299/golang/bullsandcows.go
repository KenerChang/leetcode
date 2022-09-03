package bullsandcows

import (
	"fmt"
	"strconv"
)

func getHint(secret string, guess string) string {
	// use a map to store the count of each digit
	digitMap := map[byte]int{}
	for _, digit := range secret {
		digitMap[byte(digit)]++
	}

	// count how many a
	aCount := 0
	usedMap := map[string]bool{}
	for i, digit := range guess {
		if secret[i] == guess[i] {
			digitMap[byte(digit)]--
			aCount++

			key := string(digit) + strconv.Itoa(i)
			usedMap[key] = true
		}
	}

	// count how many b
	bCount := 0
	for i, digit := range guess {
		key := string(digit) + strconv.Itoa(i)
		_, ok := usedMap[key]
		if !ok && digitMap[byte(digit)] > 0 {
			digitMap[byte(digit)]--
			bCount++
		}
	}

	return fmt.Sprintf("%dA%dB", aCount, bCount)
}
