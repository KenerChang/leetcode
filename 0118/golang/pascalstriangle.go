package pascalstriangle

func generate(numRows int) [][]int {
	result := [][]int{}
	for i := 0; i < numRows; i++ {
		for j := 0; j < i+1; j++ {
			if i == 0 {
				result = append(result, []int{1})
				continue
			}

			if j == 0 {
				result = append(result, []int{1})
			} else if j == i {
				result[i] = append(result[i], 1)
			} else {
				result[i] = append(result[i], result[i-1][j-1]+result[i-1][j])
			}
		}
	}
	return result
}
