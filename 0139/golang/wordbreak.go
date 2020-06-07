package wordbreak

func wordBreak(s string, wordDict []string) bool {
	cache := map[string]bool{}
	return wordBreakImpl(s, wordDict, cache)
}

func wordBreakImpl(s string, wordDict []string, cache map[string]bool) bool {
	if s == "" {
		return true
	}

	if breakable, found := cache[s]; found {
		return breakable
	}

	for _, word := range wordDict {
		if len(word) > len(s) || !(s[:len(word)] == word) {
			continue
		}

		remain := s[len(word):]
		breakable := wordBreakImpl(remain, wordDict, cache)
		cache[remain] = breakable
		if breakable {
			return true
		}
	}
	return false
}
