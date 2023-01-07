package ternaryexpressionparser

func parseTernary(expression string) string {
	// solve this recursively
	if len(expression) == 1 {
		return expression
	}

	// it is an express like T?5:3
	tf := expression[0]
	var separator int = 2
	var questionMarkCount int

	for {
		if expression[separator] == '?' {
			questionMarkCount++
		} else if expression[separator] == ':' {
			if questionMarkCount == 0 {
				break
			} else {
				questionMarkCount--
			}
		}
		separator++
	}

	left := parseTernary(expression[2:separator])
	right := parseTernary(expression[separator+1:])

	if tf == 'T' {
		return left
	}
	return right
}
