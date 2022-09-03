package textjustification

import (
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	lines := [][]string{}
	// for each word, if current line length + word length <= maxWidth, add it to current line
	line := []string{}
	lineCharCount := 0
	for _, word := range words {
		if len(word) >= maxWidth {
			// change line
			if len(line) > 0 {
				lines = append(lines, line)
			}

			lines = append(lines, []string{word})

			line = []string{}
			lineCharCount = 0
		} else {
			if len(word)+lineCharCount <= maxWidth {
				line = append(line, word)
				lineCharCount += len(word)

				// check if reaches maxWidth after padding a space
				if lineCharCount+1 < maxWidth {
					lineCharCount += 1
				} else {
					// change line
					lines = append(lines, line)
					line = []string{}

					lineCharCount = 0
				}

			} else {
				// create new line
				lines = append(lines, line)

				line = []string{word}

				lineCharCount = 0
				lineCharCount += len(word) + 1
			}
		}
	}

	if len(line) > 0 {
		lines = append(lines, line)
	}

	// after we get each line, do justify by line position
	results := []string{}
	for i, line := range lines {
		if i == len(lines)-1 {
			results = append(results, leftJustifyLine(line, maxWidth))
		} else {
			results = append(results, fullJustifyLine(line, maxWidth))
		}

	}
	return results
}

func fullJustifyLine(line []string, maxWidth int) string {
	// if only one word, append spaces to the end
	if len(line) == 1 {
		return leftJustifyLine(line, maxWidth)
	} else {
		return leftRightJustifyLine(line, maxWidth)
	}

}

func leftJustifyLine(line []string, maxWidth int) string {
	charCount := 0
	result := ""
	for _, word := range line {
		result += word
		charCount += len(word)

		if charCount < maxWidth {
			result += " "
			charCount += +1
		}
	}

	spaceNeeded := maxWidth - charCount
	result += strings.Repeat(" ", spaceNeeded)

	return result
}

func leftRightJustifyLine(line []string, maxWidth int) string {
	if len(line) == 2 {
		gapSize := maxWidth - len(line[0]) - len(line[1])

		return line[0] + strings.Repeat(" ", gapSize) + line[1]
	} else {
		charCount := 0
		for i, word := range line {
			charCount += len(word)
			if i != len(line)-1 {
				charCount += 1 // space
			}
		}

		gapSize := (maxWidth - charCount) / (len(line) - 1)
		unusedSpace := (maxWidth - charCount) % (len(line) - 1)

		result := ""
		for i, word := range line {
			result += word

			if i != len(line)-1 {
				result += " " + strings.Repeat(" ", gapSize)

				if unusedSpace > 0 {
					result += " "
					unusedSpace--
				}
			}
		}
		return result
	}

}
