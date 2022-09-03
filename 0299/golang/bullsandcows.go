package bullsandcows

import (
	"fmt"
)

func getHint(secret string, guess string) string {
	// use a map to store the count of each digit
	digitMap := map[byte]int{}
	aCount := 0
	bCount := 0
	for i := range secret {
		s := secret[i]
		g := guess[i]
		if s == g {
			aCount++
		} else {
			// if s is in guess, then digitMap[s] must < 0
			if digitMap[s] < 0 {
				bCount++
			}

			// i g is in secret, then digitMap[g] must > 0
			if digitMap[g] > 0 {
				bCount++
			}

			digitMap[s]++
			digitMap[g]--
		}
	}

	return fmt.Sprintf("%dA%dB", aCount, bCount)
}
