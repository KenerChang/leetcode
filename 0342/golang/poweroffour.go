package poweroffour

func isPowerOfFour(n int) bool {
	var remainder int

	for n > 1 {
		remainder = n % 4

		if remainder != 0 {
			return false
		}

		n = n / 4
	}
	return n == 1 && remainder == 0
}
