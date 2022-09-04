package minimumwindowsubstring

func minWindow(s string, t string) string {
	// if t is longer than s, return ""
	if len(t) > len(s) {
		return ""
	}

	// to solve minimum window substring problem, we need to use two pointers
	// and a map to keep how many characters we need to match

	// build t map
	tMap := map[rune]int{}
	remainMap := map[rune]int{}
	for _, char := range t {
		tMap[char]++
		remainMap[char]++
	}

	// then we use two pointers to iterate through s
	var left, right int
	var result string

	// when right pointer is moving, we check if the character is in tMap
	// it is in tMap, substract 1 from tMap
	// and if tMap becomes empty, then we find a solution
	usedmap := map[rune]int{}
	for right < len(s) {
		// if remainMap is empty, then move left pointer until tMap is not empty and left reaches a char in t
		// else move right pointer until we find a solution that remainMap is empty

		if len(remainMap) == 0 {
			// a solution candiate
			candidate := s[left:right]
			if result == "" || len(candidate) < len(result) {
				result = candidate
			}

			// move left pointer
			leftChar := rune(s[left])
			if _, ok := tMap[leftChar]; ok {
				usedmap[leftChar]--

				if usedmap[leftChar] < tMap[leftChar] {
					remainMap[leftChar]++
				}
			}
			left++
		} else {
			// move right pointer
			rightChar := rune(s[right])
			if _, ok := tMap[rightChar]; ok {
				usedmap[rightChar]++

				if usedmap[rightChar] == tMap[rightChar] {
					delete(remainMap, rightChar)
				}
			}
			right++
		}
	}

	for left < len(s) {
		if len(remainMap) == 0 {
			// a solution candiate
			candidate := s[left:right]
			if result == "" || len(candidate) < len(result) {
				result = candidate
			}

			// move left pointer
			leftChar := rune(s[left])
			if _, ok := tMap[leftChar]; ok {
				usedmap[leftChar]--

				if usedmap[leftChar] < tMap[leftChar] {
					remainMap[leftChar]++
				}
			}
		}
		left++
	}

	return result
}
