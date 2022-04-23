package maxpointsonaline

import "fmt"

func maxPoints(points [][]int) int {
	// if len(points) == 1, return 1
	if len(points) == 1 {
		return 1
	}

	// if len(points) == 2, return 2
	if len(points) == 2 {
		return 2
	}

	// iterate from first to len(points) - 2
	maxVal := 1
	for i := 0; i < len(points)-1; i++ {
		// use a map to keep if a line exist
		lines := map[string]int{}

		// for each i, check line between point i and point i+1, ...., point n
		// if line exists lines[line]++
		// else lines[line] = 1
		// fmt.Printf("points: %+v", points[i])
		for j := i + 1; j < len(points); j++ {
			line := getLine(points[i], points[j])
			// fmt.Printf("line: %s\n", line)

			if _, ok := lines[line]; !ok {
				lines[line] = 1
			}

			lines[line]++

			if lines[line] > maxVal {
				maxVal = lines[line]
			}
		}
	}
	return maxVal
}

func getLine(p1, p2 []int) string {
	// y = ax + b
	x1 := p1[0]
	y1 := p1[1]
	x2 := p2[0]
	y2 := p2[1]

	if x1 == x2 {
		return fmt.Sprintf("x=%d", x1)
	}

	numerator := y2 - y1
	denominator := x2 - x1

	// count b
	slope := float64(numerator) / float64(denominator)
	b := float64(y2) - slope*float64(x2)

	gcdXY := gcd(numerator, denominator)
	return fmt.Sprintf("%d_%d_%f", numerator/gcdXY, denominator/gcdXY, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
