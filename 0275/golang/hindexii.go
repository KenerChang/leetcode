package hindexii

func hIndex(citations []int) int {
	// if citations only has one element
	if len(citations) == 1 {
		// if has not be cited
		if citations[0] == 0 {
			return 0
		}

		// else return 1
		return 1
	}

	return hIndexRecursive(citations, 0, len(citations)-1)
}

func hIndexRecursive(citations []int, start, end int) int {
	// target: citations[i] >= n - 1 -i, find min i

	if start > end {
		return len(citations) - start
	}

	// middle point
	middle := (end + start) / 2

	// if middle equals citations[middle], return middle
	if citations[middle] == len(citations)-middle {
		return citations[middle]
	} else if citations[middle] > len(citations)-middle {
		// else if c[middle] > n - i -1, it means i is too large, find smaller i
		return hIndexRecursive(citations, start, middle-1)

	} else {
		// else c[middle] < n - i - 1, it means i is too small, find larger i
		return hIndexRecursive(citations, middle+1, end)
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
