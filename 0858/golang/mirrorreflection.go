package mirrorreflection

func gcd(a, b int) int {
	for b > 0 {
		a = a % b
		a, b = b, a
	}

	return a
}

func mirrorReflection(p int, q int) int {
	// Solve with GCD
	g := gcd(p, q)

	n := p / g
	m := q / g

	if m%2 == 0 && n%2 == 1 {
		return 0
	} else if m%2 == 1 && n%2 == 0 {
		return 2
	} else if m%2 == 1 && n%2 == 1 {
		return 1
	}

	return -1
}
