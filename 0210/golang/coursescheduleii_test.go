package coursescheduleii

import (
	"testing"
)

func verify(target, result []int) bool {
	if len(target) != len(result) {
		return false
	}

	for i, numt := range target {
		if numt != result[i] {
			return false
		}
	}
	return true
}

func TestFindOrder(t *testing.T) {
	numCourses := 2
	prerequisties := [][]int{
		[]int{1, 0},
	}

	target := []int{0, 1}
	result := findOrder(numCourses, prerequisties)
	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestFindOrderII(t *testing.T) {
	numCourses := 7
	prerequisties := [][]int{
		[]int{1, 0},
		[]int{5, 1},
		[]int{5, 2},
		[]int{3, 5},
		[]int{4, 5},
		[]int{6, 0},
	}

	target := []int{0, 1, 2, 5, 3, 4, 6}
	result := findOrder(numCourses, prerequisties)
	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestFindOrderIII(t *testing.T) {
	numCourses := 7
	prerequisties := [][]int{}

	target := []int{0, 1, 2, 3, 4, 5, 6}
	result := findOrder(numCourses, prerequisties)
	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}
