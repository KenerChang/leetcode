package dietplanperformance

func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	// init
	var kSum int
	for i := 0; i < k; i++ {
		kSum += calories[i]
	}

	result := calculatePoint(kSum, lower, upper)

	// iterate from i := 1 to n-k and count the points
	for i := 1; i+k <= len(calories); i++ {
		kSum -= calories[i-1]
		kSum += calories[i+k-1]

		result += calculatePoint(kSum, lower, upper)
	}

	return result
}

func calculatePoint(caloriesSum, lower, upper int) int {
	if caloriesSum > upper {
		return 1
	} else if caloriesSum < lower {
		return -1
	}
	return 0
}
