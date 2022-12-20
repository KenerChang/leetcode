package longestsubstringwithatmostkdistinctcharacters

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	// We can use sliding window to solve this problem

	if k == 0 {
		return 0
	}

	// use a map to keep how many distinct characters in current substring
	charsMap := map[byte]int{}

	var head, tail int
	var longest int

	for tail < len(s) {
		c := s[tail]

		// if c is in charsMap, or distinct characters count are less than k
		// put c to the map, and move tail to next
		// else remove head char's frequency in charsMap by 1 and move head to next
		if _, ok := charsMap[c]; ok || len(charsMap) < k {
			charsMap[c]++
			tail++
		} else {
			toBeRemovedChar := s[head]

			freq := charsMap[toBeRemovedChar]
			freq--

			if freq == 0 {
				delete(charsMap, toBeRemovedChar)
			} else {
				charsMap[toBeRemovedChar] = freq
			}

			head++
		}

		longest = max(longest, tail-head)
	}

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
