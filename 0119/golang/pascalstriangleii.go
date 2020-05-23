package pascalstriangleii

// import (
// 	"fmt"
// )

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	s1, s2 := make([]int, rowIndex+1), make([]int, rowIndex+1)

	var q1, q2 []int
	for i := 0; i <= rowIndex; i++ {
		if i%2 == 0 {
			q1 = s1
			q2 = s2
		} else {
			q1 = s2
			q2 = s1
		}

		for j := 0; j < i+1; j++ {
			if j == 0 {
				q1[0] = 1
			} else if j == i {
				q1[j] = 1
			} else {
				q1[j] = q2[j-1] + q2[j]
			}
		}
	}

	return q1
}
