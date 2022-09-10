package reorderedpowerof2

import (
	"math"
)

func reorderedPowerOf2(n int) bool {
	powerOf2 := math.Log2(float64(n))
	if powerOf2 == math.Ceil(powerOf2) {
		return true
	}

	// collect digits
	digitCount, digits := collectDigits(n)

	// fmt.Printf("digitCount: %d, digits: %v\n", digitCount, digits)

	// find numbers which are > 10^(x-1) and < 10^(x), and x is the number of digits in
	nums := powerOf2sBetween(math.Pow(10, float64(digitCount-1)), math.Pow(10, float64(digitCount)))
	// fmt.Printf("nums: %v\n", nums)
	for _, num := range nums {
		_, digits2 := collectDigits(num)
		// fmt.Printf("num: %d, digits2: %v\n", num, digits2)
		if compareDigits(digits, digits2) {
			return true
		}
	}

	return false
}

func collectDigits(n int) (int, map[int]int) {
	digits := map[int]int{}

	remain := n
	digitCount := 0
	for remain/10 > 0 || remain%10 > 0 {
		remainder := remain % 10
		digits[remainder]++
		digitCount++
		remain = remain / 10
	}

	return digitCount, digits
}

func compareDigits(digits1, digits2 map[int]int) bool {
	if len(digits1) != len(digits2) {
		return false
	}

	for k, v := range digits1 {
		if digits2[k] != v {
			return false
		}
	}

	for k, v := range digits2 {
		if digits1[k] != v {
			return false
		}
	}

	return true
}

func powerOf2sBetween(min, max float64) []int {
	result := []int{}

	base := math.Ceil(math.Log2(min))
	// fmt.Printf("min: %f, base: %f\n", min, base)

	for i := base; math.Pow(2, i) < max; i++ {
		result = append(result, int(math.Pow(2, i)))
	}
	return result
}
