package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	// Can use 1 pointers with memorization
	head := 0

	cache := map[byte]bool{}

	var result int
	for i, c := range s {
		// check if c is in cache
		if cache[byte(c)] {
			// if r is in cache, move head until pass last c
			for s[head] != byte(c) {

				cache[s[head]] = false
				head++
			}
			head++
		}
		// fmt.Printf("head: %d, head str: %s\n", head, string([]byte{byte(c)}))
		result = max(result, i-head+1)
		cache[byte(c)] = true
	}

	return result
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
