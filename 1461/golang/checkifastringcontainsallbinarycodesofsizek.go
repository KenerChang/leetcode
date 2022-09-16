package checkifastringcontainsallbinarycodesofsizek

func hasAllCodes(s string, k int) bool {
	// we can use a map to keep all substrings of s in length k
	// and compare if the map's length is 2^k

	substringMap := map[string]bool{}
	for i := 0; i <= len(s)-k; i++ {
		substring := s[i : i+k]
		substringMap[substring] = true
	}

	return len(substringMap) == 1<<k
}
