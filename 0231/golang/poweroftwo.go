package poweroftwo

func isPowerOfTwo(n int) bool {
	for power2 := 1; power2 <= n; power2 *= 2 {
		if power2 == n {
			return true
		}
	}
	return false
}
