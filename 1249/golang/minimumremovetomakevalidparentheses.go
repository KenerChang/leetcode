package minimumremovetomakevalidparentheses

func minRemoveToMakeValid(s string) string {
	stack := []int{}
	ignores := map[int]bool{}
	for i, c := range s {
		if c == '(' {
			// put it to the stack
			stack = append(stack, i)
		} else if c == ')' {
			// pop left parenthes
			// put the index to ignore list if there has no left parenthes
			if len(stack) == 0 {
				ignores[i] = true
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	// check if stack has unpaired left parentheses
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		ignores[top] = true
	}

	result := []byte{}
	for i := 0; i < len(s); i++ {
		if ignores[i] {
			continue
		}

		result = append(result, s[i])
	}

	return string(result)
}
