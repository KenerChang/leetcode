package fractiontorecurringdecimal

import (
	// "fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	positive := false
	if numerator == 0 {
		return "0"
	}

	if (numerator > 0 && denominator > 0) || (numerator < 0 && denominator < 0) {
		positive = true
	}
	numerator = max(numerator, -numerator)
	denominator = max(denominator, -denominator)
	quotient := numerator / denominator
	remainder := numerator % denominator

	resultStr := strconv.Itoa(quotient)
	if remainder == 0 {
		if !positive {
			resultStr = "-" + resultStr
		}
		return resultStr
	}

	remainderMap := map[int]int{}

	loop := []int{}
	var tempQ int
	var nonrepeat, repeatend string
	for remainder != 0 {
		remainder = remainder * 10
		tempQ = remainder / denominator
		if idx, found := remainderMap[remainder]; found {
			// enter loop
			nonrepeat, repeatend = loopToString(loop, idx)
			resultStr = resultStr + "." + nonrepeat + "(" + repeatend + ")"
			if !positive {
				resultStr = "-" + resultStr
			}
			return resultStr
		}
		remainderMap[remainder] = len(loop)
		// fmt.Printf("remainderMap: %+v\n", remainderMap)

		remainder = remainder % denominator
		loop = append(loop, tempQ)
	}

	// break
	nonrepeat, _ = loopToString(loop, -1)
	resultStr = resultStr + "." + nonrepeat
	if !positive {
		resultStr = "-" + resultStr
	}
	return resultStr
}

func max(num1, num2 int) int {
	if num1 >= num2 {
		return num1
	}
	return num2
}

func loopToString(loop []int, repeatIdx int) (string, string) {
	var nonrepeat, repetend string
	repeated := false
	for i, num := range loop {
		if i == repeatIdx {
			repeated = true
		}
		if !repeated {
			nonrepeat += strconv.Itoa(num)
		} else {
			repetend += strconv.Itoa(num)
		}
	}
	return nonrepeat, repetend
}
