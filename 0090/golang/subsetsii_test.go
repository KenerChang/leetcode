package subsetsii

import (
	"fmt"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	input := []int{1, 2, 2}
	target := map[string]bool{
		"[2]":     true,
		"[1]":     true,
		"[1 2 2]": true,
		"[2 2]":   true,
		"[1 2]":   true,
		"[]":      true,
	}

	result := subsetsWithDup(input)
	if !isSame(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}

	input = []int{1, 2, 2, 2}
	target = map[string]bool{
		"[2]":       true,
		"[1]":       true,
		"[1 2 2]":   true,
		"[2 2]":     true,
		"[1 2]":     true,
		"[]":        true,
		"[2 2 2]":   true,
		"[1 2 2 2]": true,
	}

	result = subsetsWithDup(input)
	if !isSame(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}

	input = []int{}
	target = map[string]bool{}

	result = subsetsWithDup(input)
	if !isSame(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}

	input = []int{1}
	target = map[string]bool{
		"[]":  true,
		"[1]": true,
	}

	result = subsetsWithDup(input)
	if !isSame(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}

	// [],[4],[3],[3,4],[3,3],[3,3,4],[2],[2,4],[2,3],[2,3,4],[2,3,3],[2,3,3,4]
	input = []int{3, 2, 3, 4}
	target = map[string]bool{
		"[]":        true,
		"[4]":       true,
		"[3]":       true,
		"[3 4]":     true,
		"[3 3]":     true,
		"[3 3 4]":   true,
		"[2]":       true,
		"[2 4]":     true,
		"[2 3]":     true,
		"[2 3 4]":   true,
		"[2 3 3]":   true,
		"[2 3 3 4]": true,
	}

	result = subsetsWithDup(input)
	if !isSame(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}

	// [[],[1],[1,4],[1,4,4],[1,4,4,4],[1,4,4,4,4],[4],[4,4],[4,4,4],[4,4,4,4]]
	input = []int{4, 4, 4, 1, 4}
	target = map[string]bool{
		"[]":          true,
		"[1]":         true,
		"[1 4]":       true,
		"[1 4 4]":     true,
		"[1 4 4 4]":   true,
		"[1 4 4 4 4]": true,
		"[4]":         true,
		"[4 4]":       true,
		"[4 4 4]":     true,
		"[4 4 4 4]":   true,
	}

	result = subsetsWithDup(input)
	if !isSame(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}

	// [[],[1],[1,4],[1,4,4],[1,4,4,4],[1,4,4,4,4],[4],[4,4],[4,4,4],[4,4,4,4]]
	input = []int{4, 4, 4}
	target = map[string]bool{
		"[]":      true,
		"[4]":     true,
		"[4 4]":   true,
		"[4 4 4]": true,
	}

	result = subsetsWithDup(input)
	if !isSame(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}

func isSame(target map[string]bool, result [][]int) bool {
	if len(target) != len(result) {
		return false
	}

	for _, nums := range result {
		key := fmt.Sprint(nums)

		if _, ok := target[key]; !ok {
			return false
		}
	}
	return true
}
