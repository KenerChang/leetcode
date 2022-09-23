package longestduplicatesubstring

func longestDupSubstring(s string) string {
	// since our goal is to find longest duplicate sub string
	// we can search solution space from size n - 1 to 1
	// and can stop searching once we find a solution

	// do binary search
	l, r := 0, len(s)-2
	var longest string
	for l <= r {
		m := (l + r) / 2

		dup := dupSubString(s, m)
		if dup != "" {
			// find search right
			longest = dup
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return longest
}

func dupSubString(s string, size int) string {
	substringMap := map[string]bool{}

	for i := 0; i+size < len(s); i++ {
		substring := s[i : i+size+1]

		if substringMap[substring] {
			return substring
		}

		substringMap[substring] = true
	}

	return ""
}
