package findthecelebrity

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		i := 0
		for i < n {
			knowAnyBody := false
			if i < n-1 {
				for j := i + 1; j < n; j++ {
					if knows(i, j) {
						i = j
						knowAnyBody = true
						break
					}
				}
			} else {
				for j := 0; j < n-1; j++ {
					if knows(i, j) {
						knowAnyBody = true
						return -1
					}
				}
			}

			if !knowAnyBody {
				// does not know any body, might be a Celebrity
				for j := 0; j < n; j++ {
					if j == i {
						continue
					}

					if !knows(j, i) {
						return -1
					}
				}

				return i
			}

		}

		return -1
	}
}
