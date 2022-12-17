package binarywatch

import "fmt"

func read(leds []int) (int, int) {
	hour := leds[0]*1 + leds[1]*2 + leds[2]*4 + leds[3]*8
	min := leds[4]*1 + leds[5]*2 + leds[6]*4 + leds[7]*8 + leds[8]*16 + leds[9]*32

	return hour, min
}

func format(hour, minute int) string {
	return fmt.Sprintf("%d:%02d", hour, minute)
}

func backtrack(leds []int, i, k int, results []string) []string {
	if k == 0 {
		h, m := read(leds)

		if (h >= 0 && h < 12) && (m >= 0 && m <= 59) {
			results = append(results, format(h, m))
		}

		return results
	}

	if i >= len(leds) {
		return results
	}

	// turn leds[i]
	leds[i] = 1
	results = backtrack(leds, i+1, k-1, results)
	leds[i] = 0

	// do not turn on leds[i]
	results = backtrack(leds, i+1, k, results)
	return results
}

func readBinaryWatch(turnedOn int) []string {
	leds := make([]int, 10)
	return backtrack(leds, 0, turnedOn, []string{})
}
