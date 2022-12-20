package longestsubstringwithatmosttwodistinctcharacters

func lengthOfLongestSubstringTwoDistinct(s string) int {
	// Use map & 2 pointers

	charMap := map[byte]int{}
	var head, tail int

	var longest int
	for tail < len(s) {
		// fmt.Printf("tail: %d, head: %d, charMap: %+v\n", tail, head, charMap)
		c := s[tail]

		if _, ok := charMap[c]; ok || len(charMap) < 2 {
			charMap[c]++

			longest = max(longest, tail-head+1)
			tail++
		} else {
			charToBeRemoved := s[head]

			freq := charMap[charToBeRemoved]
			freq--

			if freq == 0 {
				delete(charMap, charToBeRemoved)
			} else {
				charMap[charToBeRemoved] = freq
			}
			head++
		}
	}

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
