package similarrgbcolor

import "fmt"

func similarRGB(color string) string {
	// shorthand patterns: 00, 11, 22, .., FF
	shorthands := make([]int, 16)
	shorthands[1] = 17
	for i := 2; i < len(shorthands); i++ {
		shorthands[i] = shorthands[1] * i
	}

	decimalValues := parseRGB(color)

	var rShortestDistance int = decimalValues[0] * decimalValues[0]
	var rShorthandIdx int
	var gShortestDistance int = decimalValues[1] * decimalValues[1]
	var gShorthandIdx int
	var bShortestDistance int = decimalValues[2] * decimalValues[2]
	var bShorthandIdx int

	// for each shorthand, we need to compute shortest distance for R, G, B separately
	for i := 1; i < len(shorthands); i++ {
		rDistance := decimalValues[0] - shorthands[i]
		rDistance *= rDistance

		if rDistance < rShortestDistance {
			rShorthandIdx = i
			rShortestDistance = rDistance
		}

		gDistance := decimalValues[1] - shorthands[i]
		gDistance *= gDistance

		if gDistance < gShortestDistance {
			gShorthandIdx = i
			gShortestDistance = gDistance
		}

		bDistance := decimalValues[2] - shorthands[i]
		bDistance *= bDistance

		if bDistance < bShortestDistance {
			bShorthandIdx = i
			bShortestDistance = bDistance
		}
	}

	shorthandPattern := []string{
		"00", "11", "22", "33", "44", "55", "66", "77", "88", "99",
		"aa", "bb", "cc", "dd", "ee", "ff",
	}

	return fmt.Sprintf("#%s%s%s", shorthandPattern[rShorthandIdx], shorthandPattern[gShorthandIdx], shorthandPattern[bShorthandIdx])
}

func parseRGB(rgb string) []int {
	hexDigits := map[byte]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5,
		'6': 6, '7': 7, '8': 8, '9': 9, 'a': 10, 'b': 11,
		'c': 12, 'd': 13, 'e': 14, 'f': 15,
	}

	result := make([]int, 3)
	result[0] = hexDigits[rgb[1]]*16 + hexDigits[rgb[2]]
	result[1] = hexDigits[rgb[3]]*16 + hexDigits[rgb[4]]
	result[2] = hexDigits[rgb[5]]*16 + hexDigits[rgb[6]]

	return result
}
