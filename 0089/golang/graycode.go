package graycode

func grayCode(n int) []int {
	// gray codes is composed of gray codes of n-1 and reverse of gray codes of n-1 add add 1 as first bit
	if n == 0 {
		return []int{0}
	}

	if n == 1 {
		return []int{0, 1}
	}

	grayCodes := grayCode(n - 1)
	count := len(grayCodes) - 1

	prefix := 1 << uint(n-1) // add 1 as first bit
	for count >= 0 {
		grayCodes = append(grayCodes, grayCodes[count]+prefix)
		count--
	}
	return grayCodes
}
