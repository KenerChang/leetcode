package adddigits

func addDigits(num int) int {
	if num >= 0 && num < 10 {
		return num
	}
	digits := toDigits(num)

	for len(digits) > 1 {
		last := len(digits) - 1
		last2 := last - 1

		ds := toDigits(digits[last] + digits[last2])
		// fmt.Printf("ds: %v, last: %d, last2: %d\n", ds, digits[last], digits[last2])

		digits = append(digits[0:last2], ds...)
		// fmt.Printf("digits: %v\n", digits)
	}
	return digits[0]
}

func toDigits(num int) []int {
	result := []int{}

	for num > 0 {
		quotient := num / 10
		remainder := num % 10

		result = append(result, remainder)

		num = quotient
	}

	return result
}
