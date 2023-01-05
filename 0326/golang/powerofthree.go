package powerofthree

func isPowerOfThree(n int) bool {
	var remain int
	for n > 1 {
		remain = n % 3
		if remain != 0 {
			return false
		}

		n = n / 3
	}
	return n == 1 && remain == 0
}
