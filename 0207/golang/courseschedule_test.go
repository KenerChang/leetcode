package courseschedule

import (
	"testing"
)

// func TestCanFinish(t *testing.T) {
// 	numCourses := 2
// 	prerequisites := [][]int{
// 		[]int{1, 0},
// 	}
// 	target := true

// 	result := canFinish(numCourses, prerequisites)
// 	if target != result {
// 		t.Errorf("expect %t, got %t", target, result)
// 	}
// }

// func TestCanFinishII(t *testing.T) {
// 	numCourses := 2
// 	prerequisites := [][]int{
// 		[]int{1, 0},
// 		[]int{0, 1},
// 	}
// 	target := false

// 	result := canFinish(numCourses, prerequisites)
// 	if target != result {
// 		t.Errorf("expect %t, got %t", target, result)
// 	}
// }

// func TestCanFinishIII(t *testing.T) {
// 	numCourses := 3
// 	prerequisites := [][]int{
// 		[]int{0, 1},
// 		[]int{1, 2},
// 		[]int{2, 0},
// 	}
// 	target := false

// 	result := canFinish(numCourses, prerequisites)
// 	if target != result {
// 		t.Errorf("expect %t, got %t", target, result)
// 	}
// }

func TestCanFinishIV(t *testing.T) {
	numCourses := 8
	prerequisites := [][]int{
		[]int{1, 0},
		[]int{2, 6},
		[]int{1, 7},
		[]int{6, 4},
		[]int{7, 0},
		[]int{0, 5},
	}
	target := true

	result := canFinish(numCourses, prerequisites)
	if target != result {
		t.Errorf("expect %t, got %t", target, result)
	}
}
