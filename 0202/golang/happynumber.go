package happynumber

// import (
// 	"fmt"
// )

func isHappy(n int) bool {
	var next int
	cache := map[int]bool{}

	for true {
		if _, found := cache[n]; found {
			return false
		}

		if n == 0 {
			if next == 1 {
				return true
			}
			cache[next] = true
			n = next
			next = 0
		}

		remainder := n % 10
		next += remainder * remainder
		n = n / 10
	}

	return false
}
