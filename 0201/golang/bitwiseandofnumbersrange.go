package bitwiseandofnumbersrange

func rangeBitwiseAnd(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if m == n {
		return m
	}

	diff := n - m
	maxPowerOfDiff := getMaxPowerOfTwo(diff) + 1
	n = n >> maxPowerOfDiff << maxPowerOfDiff
	m = m >> maxPowerOfDiff << maxPowerOfDiff

	return n & m
}

func getMaxPowerOfTwo(num int) int {
	power := 0
	for num >= 2 {
		power++
		num = num >> 1
	}
	return power
}
