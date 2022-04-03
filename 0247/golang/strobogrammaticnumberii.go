package strobogrammaticnumberii

var (
	pairs [][]byte = [][]byte{
		{'0', '0'},
		{'1', '1'},
		{'6', '9'},
		{'8', '8'},
		{'9', '6'},
	}
)

func findStrobogrammatic(n int) []string {
	return findStrobogrammaticRecursive(-1, n, make([]byte, n), []string{})
}

func findStrobogrammaticRecursive(ptrHead, ptrTail int, proceeded []byte, results []string) []string {
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

		results = findStrobogrammaticRecursive(ptrHead, ptrTail, proceeded, results)
	}

	return results
}
