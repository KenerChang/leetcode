package wordpatternii

func recursive(pattern string, s string, patternMap map[byte]string, patternExistMap map[string]byte) bool {
	if len(pattern) == 0 || len(s) == 0 {
		if len(pattern) == 0 && len(s) == 0 {
			return true
		}
		return false
	}

	if len(pattern) == 1 {
		if pat, ok := patternMap[pattern[0]]; ok {
			return pat == s
		} else {
			if char, ok := patternExistMap[s]; ok {
				return char == pattern[0]
			}
		}
	}

	for i := range pattern {
		char := pattern[i]

		if pat, ok := patternMap[char]; ok {
			if len(s) < len(pat) || pat != s[0:len(pat)] {
				return false
			} else {
				return recursive(pattern[i+1:], s[len(pat):], patternMap, patternExistMap)
			}
		}

		for j := range s {
			pat := s[0 : j+1]

			if _, ok := patternExistMap[pat]; ok {
				continue
			}

			patternMap[char] = pat
			patternExistMap[pat] = char

			matched := recursive(pattern[i+1:], s[j+1:], patternMap, patternExistMap)
			if matched {
				return true
			}

			delete(patternMap, char)
			delete(patternExistMap, pat)
		}
		return false
	}

	return false
}

func wordPatternMatch(pattern string, s string) bool {
	return recursive(pattern, s, map[byte]string{}, map[string]byte{})
}
