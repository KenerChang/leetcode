package additivenumber

import (
	"strconv"
)

func recursive(num string, prev1, prev2, total int) bool {
	if len(num) == 0 {
		return total >= 3
	}

	if prev1 == -1 || prev2 == -1 {
		for i := range num {
			if num[0] == '0' && i > 0 {
				return false
			}

			if prev1 == -1 {
				prev1, _ = strconv.Atoi(num[:i+1])
				isAdditive := recursive(num[i+1:], prev1, -1, 1)

				if isAdditive {
					return true
				}

				prev1 = -1
			} else if prev2 == -1 {
				prev2, _ = strconv.Atoi(num[:i+1])
				isAdditive := recursive(num[i+1:], prev1, prev2, 2)

				if isAdditive {
					return true
				}

				prev2 = -1
			}
		}
	}

	if prev1 == -1 || prev2 == -1 {
		return false
	}

	prevSum := prev1 + prev2
	for i := range num {
		if num[0] == '0' && i > 0 {
			return false
		}

		curNum, _ := strconv.Atoi(num[:i+1])
		if curNum == prevSum {
			return recursive(num[i+1:], prev2, curNum, 3)
		} else if curNum > prevSum {
			return false
		}
	}

	return false
}

func isAdditiveNumber(num string) bool {
	// Fulfill 3 requirements
	// 1. at least 3 nums
	// 2. no leading 0s
	// 3. is additive sequence
	return recursive(num, -1, -1, -1)
}
