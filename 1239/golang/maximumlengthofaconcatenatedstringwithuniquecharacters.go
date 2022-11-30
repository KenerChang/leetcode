package maximumlengthofaconcatenatedstringwithuniquecharacters

func dfs(arr []int, pos int, result int) int {
	// check if chars in result are duplicated
	resultChars := result & ((1 << 26) - 1)
	resultLen := result >> 26

	best := resultLen
	for i := pos; i < len(arr); i++ {
		currChars := arr[i] & ((1 << 26) - 1)
		currLen := arr[i] >> 26

		// check if chars duplicate
		if currChars&resultChars > 0 {
			continue
		}

		newResult := currChars + resultChars + ((currLen + resultLen) << 26)
		best = max(best, dfs(arr, i+1, newResult))

	}
	return best
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxLength(arr []string) int {
	// transform word to int
	bitmaps := []int{}
	for _, s := range arr {
		bm := wordToBitmap(s)

		if bm == -1 {
			continue
		}

		bitmaps = append(bitmaps, bm)
	}

	// do dfs
	return dfs(bitmaps, 0, 0)
}

func wordToBitmap(s string) int {
	var result int
	for _, c := range s {
		bitPos := 1 << (c - 'a')

		if result&bitPos > 0 {
			return -1
		}

		result += bitPos
	}

	result = result + (len(s) << 26)

	return result
}
