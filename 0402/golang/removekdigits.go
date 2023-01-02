package removekdigits

func removeKdigits(num string, k int) string {
	stack := []byte{}

	for i := range num {
		c := num[i]

		for len(stack) > 0 && c < stack[len(stack)-1] && k > 0 {
			stack = stack[0 : len(stack)-1]
			k--
		}

		if c == '0' {
			if len(stack) > 0 {
				stack = append(stack, c)
			}
		} else {
			stack = append(stack, c)
		}
	}

	if k >= len(stack) {
		return "0"
	} else if len(stack) > k {
		stack = stack[0 : len(stack)-k]
	}

	return string(stack)
}
