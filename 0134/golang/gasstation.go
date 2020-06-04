package gasstation

func canCompleteCircuit(gas []int, cost []int) int {
	size := len(gas)
	accumlated := 0
	consumption := 0
	candidates := []int{}
	for i, fill := range gas {
		accumlated += fill
		consumption += cost[i]

		if fill >= cost[i] {
			candidates = append(candidates, i)
		}
	}

	if accumlated < consumption {
		return -1
	}

	for _, candidate := range candidates {
		if canComplete(candidate, size, gas, cost) {
			return candidate
		}
	}

	return -1
}

func canComplete(startFrom, size int, gas []int, cost []int) bool {
	accumlated := 0
	for i := startFrom; i < startFrom+size; i++ {
		idx := i % size
		accumlated += gas[idx]
		accumlated -= cost[idx]

		if accumlated < 0 {
			return false
		}
	}
	return true
}
