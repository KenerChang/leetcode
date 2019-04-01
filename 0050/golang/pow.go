package pow

// import "fmt"

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	if n == 0 || x == 1 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n == -1 {
		return float64(1) / x
	}

	countN := n // absoulte n for count
	if countN < 0 {
		countN = -countN
	}

	power := 1
	num := x
	previous := x

	// power left shift, num multiply x
	for power < countN {
		power = power << 1
		previous = num
		num = num * num
	}
	power = power >> 1
	remianPow := countN - power

	num = previous * myPow(x, remianPow)

	if n > 0 {
		return num
	} else {
		return float64(1) / num
	}
}
