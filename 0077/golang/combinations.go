package combinations

// import "fmt"

// combine
func combine(n int, k int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	if k == 0 {
		return [][]int{[]int{}}
	}

	if k == 1 {
		result := [][]int{}
		for i := 1; i <= n; i++ {
			result = append(result, []int{i})
		}
		return result
	}

	if n == k {
		result := []int{}
		for i := 1; i <= n; i++ {
			result = append(result, i)
		}
		return [][]int{result}
	}

	setsOFEachNum := map[int][][]int{}

	for i := 1; i <= k; i++ {
		setsOFEachNum[i] = [][]int{}
	}

	for i := 1; i <= n; i++ {
		if _, ok := setsOFEachNum[1]; !ok {
			setsOFEachNum[1] = [][]int{}
		}
		var previousCount int
		for j := 1; j <= k; j++ {
			sets, _ := setsOFEachNum[j-1]
			if j == 1 {
				setsOFEachNum[1] = append(setsOFEachNum[1], []int{i})
				previousCount = 1
				continue
			}

			for s := 0; s < len(sets)-previousCount; s++ {
				set := sets[s]
				newSet := append([]int{}, set...)
				newSet = append(newSet, i)

				setsOFEachNum[j] = append(setsOFEachNum[j], newSet)
			}
			previousCount = len(sets) - previousCount
		}
	}
	return setsOFEachNum[k]
}
