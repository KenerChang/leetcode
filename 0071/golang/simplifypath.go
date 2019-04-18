package simplifypath

import "fmt"
import "strings"

func simplifyPath(path string) string {
	tokens := strings.Split(path, "/")

	stack := []string{}
	for _, token := range tokens {
		switch token {
		case ".":
			break
		case "":
			break
		case " ":
			break
		case "..":
			// pop stack
			if len(stack) > 0 {
				stack = stack[0 : len(stack)-1]
			}
		default:
			stack = append(stack, token)
		}
	}

	result := ""
	for _, token := range stack {
		result = fmt.Sprintf("%s/%s", result, token)
	}

	if result == "" {
		result = "/"
	}

	return result
}
