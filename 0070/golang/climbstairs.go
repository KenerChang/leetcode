package climbstairs

func climbStairs(n int) int {
	// use DP to solve this problem
	// f(x) = f(x-1) + f(x-2) // if x >2
	// f(x) = 1 // if x == 1
	// f(x) = 2 // if x == 2

	if n == 0 {
		return 1
	}

	if n == 1 || n == 2 {
		return n
	}

	cache := make([]int, n)
	cache[0] = 1 // zero-based
	cache[1] = 2 //

	for i := 2; i < n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}

	return cache[n-1]
}
