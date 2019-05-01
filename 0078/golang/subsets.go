package subsets

// subsets create power sets of given numbers
func subsets(nums []int) [][]int {
	result := [][]int{
		[]int{},
	}

	// each number is either in subset or not in subset
	// so copying existed subsets and adding the number to subsets means the number in subsets
	for _, num := range nums {
		for _, subset := range result {
			newSet := append([]int{}, subset...)
			newSet = append(newSet, num)

			result = append(result, newSet)
		}
	}
	return result
}
