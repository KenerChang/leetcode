package firstbadversion

var isBadVersion func(version int) bool

func firstBadVersion(n int) int {
	// binary search

	return firstBadVersionRecursive(1, n)
}

func firstBadVersionRecursive(start, end int) int {
	// binary search

	// if start >= end, we found return start
	if start >= end {
		return start
	}

	// check if middle is bad version
	middle := (end + start) / 2

	// if middle is a bad version, check start ~ middle
	// else middle is not a bad version, check middle +1 ~ end
	if isBadVersion(middle) {
		return firstBadVersionRecursive(start, middle)
	} else {
		return firstBadVersionRecursive(middle+1, end)
	}
}
