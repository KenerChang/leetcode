package regularexpressionmatching

func isMatch(s string, p string) bool {
	if p == ".*" {
		return true
	}

	if s == "" && p == "" {
		return true
	}

	var sPtr, pPtr int
	// iterate p & s by pPtr & sPtr
	for pPtr < len(p) {
		var nextIsStar bool
		// check next char of p
		if pPtr < len(p)-1 {
			nextIsStar = p[pPtr+1] == '*'
		}

		// if p[pPtr+1] is not *, check if s[sPtr] != p[Ptr]
		if !nextIsStar {
			if sPtr >= len(s) {
				return false
			}

			if s[sPtr] != p[pPtr] && p[pPtr] != '.' {
				return false
			}

			sPtr++
			pPtr++
		} else {
			// else
			if sPtr >= len(s) {
				pPtr += 2
			} else if p[pPtr] == '.' || s[sPtr] == p[pPtr] {
				// case 2: p[ptr] is not "." and s[sPtr] == p[Ptr] return isMatch(s[sPtr+1:], p[pPtr:]) || isMatch(s[sPtr:], p[pPtr+2:])
				// case 3: p[ptr] is ".", return isMatch(s[sPtr+1:], p[pPtr:]) || isMatch(s[sPtr:], p[pPtr+2:])
				result := isMatch(s[sPtr+1:], p[pPtr:]) || isMatch(s[sPtr:], p[pPtr+2:])
				return result
			} else {
				// case 1: p[Ptr] is not "." and s[sPtr] != p[Ptr], move to next
				pPtr += 2
			}
		}
	}

	return sPtr >= len(s)
}
