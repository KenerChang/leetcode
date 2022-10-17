package uniquenumberofoccurrences

func uniqueOccurrences(arr []int) bool {
	// We can solve this problem with  hashmapes

	// first use a hashmap to keep num count
	numCounts := map[int]int{}
	for _, num := range arr {
		numCounts[num]++
	}

	// then use hashmap to detect if conflict
	countMap := map[int]bool{}
	for _, count := range numCounts {
		if countMap[count] {
			return false
		}

		countMap[count] = true
	}

	return true
}
