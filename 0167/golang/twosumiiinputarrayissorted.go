package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	var sum int
	for true {
		sum = numbers[i] + numbers[j]
		if sum > target {
			j--
		} else if sum < target {
			i++
		} else {
			return []int{i + 1, j + 1}
		}
	}
	return []int{i + 1, j + 1}
}
