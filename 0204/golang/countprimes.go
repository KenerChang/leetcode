package countprimes

import (
// "math"
)

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	isPrime := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if !isPrime[i] {
			count++
			for j := 2; i*j < n; j++ {
				isPrime[i*j] = true
			}
		}
	}
	return count
}
