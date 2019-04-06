package uniquepaths

// import "fmt"

func uniquePaths(m int, n int) int {
	// the asnwer formula is combination m-1 of n-1
	if m == 0 || n == 0 {
		return 0
	}

	if m == 1 || n == 1 {
		return 1
	}

	realM := m - 1
	realN := n - 1
	total := realM + realN
	smaller := Min(realM, realN)
	larger := total - smaller
	smallerFictoral := 1
	mnfictoral := 1

	for i := 2; i <= total; i++ {
		if i <= smaller {
			smallerFictoral *= i
		} else if i > larger {
			mnfictoral *= i
		}
	}

	return mnfictoral / smallerFictoral
}

func Min(i1, i2 int) int {
	if i1 <= i2 {
		return i1
	} else {
		return i2
	}
}
