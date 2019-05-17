package decodeways

import "strconv"

func numDecodings(s string) int {
	// dynamic programming
	// ways[n] = ways[n-1] + ways[n-2] if s[n-1][n] exist
	// wyas[n] = ways[n-1] if s[n-1][n] does not exist

	stringLen := len(s)
	if stringLen == 0 {
		return stringLen
	}

	ways := make([]int, stringLen)
	ways[0] = 0
	if isNum(s[0:1]) {
		ways[0] = 1
	}

	if stringLen == 1 {
		return ways[0]
	}

	for i := 1; i < stringLen; i++ {
		ways[i] = ways[i-1]
		if s[i] == '0' {
			ways[i] = 0
		}

		if isNum(s[i-1 : i+1]) {
			if i > 1 {
				ways[i] += ways[i-2]
			} else {
				ways[i]++
			}
		}
	}

	return ways[stringLen-1]
}

func isNum(s string) bool {
	if s[0] == '0' {
		return false
	}

	num, _ := strconv.Atoi(s)
	if num <= 26 && num > 0 {
		return true
	} else {
		return false
	}
}
