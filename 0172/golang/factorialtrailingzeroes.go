package factorialtrailingzeroes

func trailingZeroes(n int) int {
	base := 5

	var count, quotient int
	for n/base > 0 {
		quotient = n / base
		count += quotient

		base *= 5
	}
	return count
}
