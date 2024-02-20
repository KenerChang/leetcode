package strobogrammaticnumberiii

import (
	"strconv"
)

var (
	pairs [][]byte = [][]byte{
		{'0', '0'},
		{'1', '1'},
		{'6', '9'},
		{'8', '8'},
		{'9', '6'},
	}
)

func strobogrammaticInRange(low string, high string) int {
	lowInt, _ := strconv.ParseInt(low, 10, 64)
	highInt, _ := strconv.ParseInt(high, 10, 64)

	var result int
	for n := len(low); n <= len(high); n++ {
		nums := backtracking(-1, int64(n), make([]byte, n), []string{})

		for _, num := range nums {
			numInt, _ := strconv.ParseInt(num, 10, 64)
			if numInt >= lowInt && numInt <= highInt {
				result++
			}
		}
	}
	return result
}

func backtracking(ptrHead, ptrTail int64, proceeded []byte, results []string) []string {
	// candidate pairs: (0,0), (1,1), (6,9), (8,8), (9,6)

	// if match top condition, do generate results
	if ptrHead >= ptrTail-1 {
		results = append(results, string(proceeded))
		return results
	}

	// put a candidate pair and proceed to next position
	ptrHead++
	ptrTail--
	for _, pair := range pairs {
		// assign

		// ptrHead == ptrTail, check if the pair can be used
		if ptrHead == ptrTail {
			// can only use 0, 1, 8
			if pair[0] == '0' || pair[0] == '1' || pair[0] == '8' {
				proceeded[ptrHead] = pair[0]
			} else {
				continue
			}
		} else if ptrHead == 0 && pair[0] == '0' {
			continue
		} else {
			// ptrHead < ptrTail
			proceeded[ptrHead] = pair[0]
			proceeded[ptrTail] = pair[1]
		}

		results = backtracking(ptrHead, ptrTail, proceeded, results)
	}

	return results
}
