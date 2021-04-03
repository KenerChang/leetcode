package basiccalculatorii

import (
	"strconv"
	"strings"
)

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	if len(s) == 0 {
		return 0
	}
	start, end := 0, 1
	prev, result := 0, 0

	for end <= len(s) {
		if end == len(s) || isOperator(s[end]) {
			num, _ := strconv.Atoi(s[start:end])
			if start == 0 {
				result += num
				prev = result
			} else {
				operator := s[start-1] // previous operator
				switch operator {
				case '+':
					result += num
					prev = num
				case '-':
					result -= num
					prev = -num
				case '*':
					result = result - prev + prev*num
					prev = prev * num
				case '/':
					result = result - prev + prev/num
					prev = prev / num
				}
			}
			start = end + 1
		}
		end++
	}
	return result
}

func isOperator(char byte) bool {
	if char == '-' || char == '+' || char == '/' || char == '*' {
		return true
	}
	return false
}
