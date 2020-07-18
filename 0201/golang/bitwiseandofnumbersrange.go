package bitwiseandofnumbersrange

func rangeBitwiseAnd(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if m == n {
		return m
	}

	maxPowerOfM := getMaxPowerOfTwo(m)
	maxPowerOfN := getMaxPowerOfTwo(n)

	if maxPowerOfM != maxPowerOfN {
		return 0
	}

	maxPowerOfDiff := getMaxPowerOfTwo(n - m)
	result := 1 << maxPowerOfM
	for power := maxPowerOfM - 1; power > maxPowerOfDiff; power-- {
		divider := 1 << power
		if m&divider > 0 && n&divider > 0 {
			result += divider
		}
	}
	return result
}

func getMaxPowerOfTwo(num int) int {
	maxPower := 0
	power := 0
	for num > 0 {
		if num%2 > 0 {
			maxPower = power
		}
		power++
		num = num / 2
	}
	return maxPower
}
