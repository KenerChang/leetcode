package gasstation

func canCompleteCircuit(gas []int, cost []int) int {
	size := len(gas)
	idx := 0
	sum := 0

	for idx < size {
		currIdx := idx
		thisIdxSum := 0

		for idx < size {
			thisIdxSum += gas[idx] - cost[idx]
			idx++

			if thisIdxSum < 0 {
				// condition break
				break
			}
		}

		sum += thisIdxSum
		if sum < 0 {
			// sum of 0 to i-1 was negative
			// keep test if i can complete
			continue
		}
		return currIdx
	}
	return -1
}
