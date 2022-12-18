package threesumwithmultiplicity

var (
	m = int(1e9 + 7)
)

func twoSumMutli(arr []int, target int) int {
	freqMap := map[int]int{}

	var result int
	for _, num := range arr {
		complement := target - num
		if freq, ok := freqMap[complement]; ok {
			result = (result + freq) % m
		}

		freqMap[num]++
	}

	return result
}

func threeSumMulti(arr []int, target int) int {
	var result int
	for i, num := range arr {
		complement := target - num

		pairs := twoSumMutli(arr[i+1:], complement)

		result += pairs
		result = result % m
	}
	return result
}
