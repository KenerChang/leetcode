package countprimes

import (
	"math"
)

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	primes := []int{2}
	for i := 3; i < n; i++ {
		if !isPrime(i, primes) {
			continue
		}

		primes = append(primes, i)
	}
	return len(primes)
}

func isPrime(n int, primes []int) bool {
	sr := math.Sqrt(float64(n))
	for _, prime := range primes {
		if float64(prime) > sr {
			break
		}
		if n%prime == 0 {
			return false
		}
	}
	return true
}
