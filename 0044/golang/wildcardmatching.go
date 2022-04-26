package wildcardmatching

import "fmt"

func isMatch(s string, p string) bool {
	return isMatchRecursive(s, p, map[string]bool{})
}

func isMatchRecursive(s string, p string, cache map[string]bool) bool {
	// To solve if a pattern matches a string, our target is
	// iterating both s and p, and we expect that we can reach
	// the ends of both s and p.

	key := fmt.Sprintf("%s_%s", s, p)

	if result, ok := cache[key]; ok {
		return result
	}

	if p == "*" {
		cache[key] = true
		return true
	}

	var sPtr, pPtr int
	for sPtr < len(s) && pPtr < len(p) {
		// case 1: p[pPtr] == '?', both sPtr and pPtr move 1 postion
		switch p[pPtr] {
		case '?':
			// case 1: p[pPtr] == '?', both sPtr and pPtr move 1 postion
			sPtr++
			pPtr++
		case '*':
			// case 2: p[pPtr] == '*', return or of the following 3 case
			// case 2.1 move s for 1 position
			// case 2.2 move s and p for 1 position
			// case 2.3 move p for 1 position

			// if any case is matched, then s matches p
			result := isMatchRecursive(s[sPtr+1:], p[pPtr:], cache) || isMatchRecursive(s[sPtr:], p[pPtr+1:], cache) || isMatchRecursive(s[sPtr+1:], p[pPtr+1:], cache)
			cache[key] = result
			return result
		default:
			// case 3: p[pPtr] != '?' && p[pPtr] != '*',
			// check if s[sPtr] == p[pPtr], it the two char matches, move to next postiion
			// else s does not match p
			if s[sPtr] == p[pPtr] {
				sPtr++
				pPtr++
			} else {
				cache[key] = false
				return false
			}
		}
	}

	// p reaches end, but s does not
	if sPtr < len(s) {
		cache[key] = false
		return false
	}

	// s reach end, but p does not
	// check if p contains characters other than '?'
	for pPtr < len(p) {
		if p[pPtr] != '*' {
			cache[key] = false
			return false
		}

		pPtr++
	}

	cache[key] = true
	return true
}
