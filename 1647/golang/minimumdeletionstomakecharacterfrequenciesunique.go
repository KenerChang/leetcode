package minimumdeletionstomakecharacterfrequenciesunique

import (
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func minDeletions(s string) int {
	// first, get character frequencies
	freqs := make([]int, 26)
	for _, char := range s {
		freqs[char-'a']++
	}

	// sort by frequencies
	sort.Ints(freqs)

	// for each freq, if freq > maxAllowed => result += freq - maxAllowed
	// maxAllowed = freq - 1
	var result int
	maxAllowed := freqs[len(freqs)-1]
	for i := len(freqs) - 1; i >= 0; i-- {
		if freqs[i] > maxAllowed {
			result += freqs[i] - maxAllowed
			freqs[i] = maxAllowed
		}
		maxAllowed = max(0, freqs[i]-1)
	}

	return result
}
