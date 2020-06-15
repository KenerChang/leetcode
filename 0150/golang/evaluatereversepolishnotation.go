package evaluatereversepolishnotation

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		if token == "*" {
			last2, last := len(stack)-2, len(stack)-1
			stack[last2] = stack[last2] * stack[last]
			stack = stack[0:last]
		} else if token == "/" {
			last2, last := len(stack)-2, len(stack)-1
			stack[last2] = stack[last2] / stack[last]
			stack = stack[0:last]
		} else if token == "+" {
			last2, last := len(stack)-2, len(stack)-1
			stack[last2] = stack[last2] + stack[last]
			stack = stack[0:last]
		} else if token == "-" {
			last2, last := len(stack)-2, len(stack)-1
			stack[last2] = stack[last2] - stack[last]
			stack = stack[0:last]
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}
