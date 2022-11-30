package concatenationofconsecutivebinarynumbers

var (
	m int = 10e8 + 7
)

func concatenatedBinary(n int) int {
	var result int
	var digits int = 0

	for i := 1; i <= n; i++ {
		if (i & (i - 1)) == 0 {
			digits++
		}

		// mod m
		// fmt.Printf("result: %d, count: %d, i: %d, %d\n", result, count, i, ((result << count) + i))
		result = update(result, digits, i)
	}
	return result
}

func update(result, count, i int) int {
	countPart := 1 << count
	iPart := i % m

	result = (result * countPart) % m
	result = (result + iPart) % m

	return result
}
