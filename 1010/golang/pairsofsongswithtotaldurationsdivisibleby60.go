package pairsofsongswithtotaldurationsdivisibleby60

func numPairsDivisibleBy60(time []int) int {
	// We can use a map to keep remainder of time and its position
	// and for each time, we can find its complementary remainder
	// and we only positions that behind this number

	remainderMap := map[int][]int{}
	for i, duration := range time {
		remainder := duration % 60
		if _, ok := remainderMap[remainder]; !ok {
			remainderMap[remainder] = []int{}
		}

		remainderMap[remainder] = append(remainderMap[remainder], i)
	}

	var result int
	for i, duration := range time {
		remainder := duration % 60
		c := 60 - remainder
		c = c % 60

		positions, ok := remainderMap[c]
		if !ok {
			continue
		}

		for _, pos := range positions {
			if pos > i {
				result++
			}
		}
	}
	return result
}
