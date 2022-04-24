package perfectsquares

import (
	"math"
)

func numSquares(n int) int {
	// for i from 1 to Sqrt(n)
	//
	return numSquaresRecursive(n, map[int]int{})
}

func numSquaresRecursive(n int, cache map[int]int) int {
	if n == 1 {
		return 1
	}

	// if n in cache, return result
	if result, ok := cache[n]; ok {
		return result
	}

	sqrt := math.Floor(math.Sqrt(float64(n)))
	// fmt.Printf("n: %d, sqrt: %f\n", n, sqrt)

	if int(sqrt) == 1 {
		return n
	}

	// for Sqrt(n) to 1
	minVal := math.MaxInt
	for i := int(sqrt); i > 0; i-- {
		// result = 1 + numSquaresRecursive(n - i*i)
		var result int
		remain := n - i*i
		if remain == 0 {
			result = 1
		} else {
			result = 1 + numSquaresRecursive(n-i*i, cache)
		}

		// if result < minVal, minVal
		if result < minVal {
			minVal = result
		}
	}

	cache[n] = minVal
	// fmt.Printf("cache: %+v\n", cache)
	return minVal
}
