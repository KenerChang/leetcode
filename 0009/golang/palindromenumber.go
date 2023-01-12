package palindromenumber

func isPalindrome(x int) bool {
	// return fales if x is less than 0

	// reverse x and store to rerv
	var rerv int
	n := x
	for n > 0 {
		remainder := n % 10

		rerv = rerv * 10
		rerv += remainder

		n = n / 10
	}

	return rerv == x
}
