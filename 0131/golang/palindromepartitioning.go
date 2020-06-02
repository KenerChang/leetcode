package palindromepartitioning

func partition(s string) [][]string {
	var curr []string
	var res [][]string
	res = partitionImpl(s, 0, res, curr)
	return res
}

func partitionImpl(s string, start int, result [][]string, current []string) [][]string {
	strLen := len(s)
	if start == strLen {
		// reach end, add current to result
		currCopy := make([]string, len(current))
		copy(currCopy, current)
		result = append(result, currCopy)
		return result
	}

	for i := start; i < strLen; i++ {
		if !isPalindrome(s, start, i) {
			continue
		}

		current = append(current, s[start:i+1])
		result = partitionImpl(s, i+1, result, current)
		current = current[:len(current)-1]
	}
	return result
}

func isPalindrome(s string, start int, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
