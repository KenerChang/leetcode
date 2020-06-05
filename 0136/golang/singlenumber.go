package singlenumber

func singleNumber(nums []int) int {
	visited := map[int]bool{}
	result := 0

	for _, num := range nums {
		if _, appear := visited[num]; appear {
			result -= num
		} else {
			result += num
			visited[num] = true
		}
	}
	return result
}
