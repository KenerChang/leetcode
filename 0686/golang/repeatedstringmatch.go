package repeatedstringmatch

import (
	"math"
	"strings"
)

func repeatedStringMatch(a string, b string) int {
	// if x times of a can contains b, then x must be floor(len(b)/len(a)) or ceil(len(b)/len(a))
	times := float64(len(b)) / float64(len(a))
	ceil := math.Ceil(times)

	if strings.Contains(strings.Repeat(a, int(ceil)), b) {
		return int(ceil)
	}

	if strings.Contains(strings.Repeat(a, int(ceil)+1), b) {
		return int(ceil) + 1
	}

	return -1
}
