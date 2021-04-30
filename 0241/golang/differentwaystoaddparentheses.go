package differentwaystoaddparentheses

import "strconv"

func diffWaysToCompute(expression string) []int {
	// return diffWaysToComputeRecursive(expression, map[string][]int{})
	num, err := strconv.Atoi(expression)
	if err == nil {
		return []int{num}
	}

	results := []int{}
	for i, char := range expression {
		if char == 45 && i == 0 {
			continue
		}

		if char >= 48 && char <= 57 {
			continue
		}

		lefts := diffWaysToCompute(expression[0:i])
		rights := diffWaysToCompute(expression[i+1:])

		for _, left := range lefts {
			for _, right := range rights {
				num := doOperation(left, right, char)
				results = append(results, num)
			}
		}
	}
	return results
}

func doOperation(leftOP, rightOP int, operator rune) int {
	switch operator {
	case 43:
		// +
		return leftOP + rightOP
	case 45:
		// -
		return leftOP - rightOP
	case 42:
		// *
		return leftOP * rightOP
	}
	return 0
}
