package findmodeinbinarysearchtree

import (
	"testing"
)

func verify(target, result []int) bool {
	if len(target) != len(result) {
		return false
	}

	targetMap := map[int]bool{}
	resultMap := map[int]bool{}

	for i := 0; i < len(target); i++ {
		targetNum := target[i]
		resultNum := result[i]

		targetMap[targetNum] = true
		resultMap[resultNum] = true
	}

	for key := range targetMap {
		if _, found := resultMap[key]; !found {
			return false
		}
	}

	for key := range resultMap {
		if _, found := targetMap[key]; !found {
			return false
		}
	}
	return true
}

// func TestFindModeI(t *testing.T) {
// 	tree := &TreeNode{
// 		Val: 1,
// 		Right: &TreeNode{
// 			Val: 2,
// 			Left: &TreeNode{
// 				Val: 2,
// 			},
// 		},
// 	}

// 	target := []int{2}
// 	result := findMode(tree)
// 	if !verify(target, result) {
// 		t.Errorf("expect %+v, got: %+v", target, result)
// 	}
// }

// func TestFindModeII(t *testing.T) {
// 	tree := &TreeNode{
// 		Val: 1,
// 		Right: &TreeNode{
// 			Val: 2,
// 		},
// 	}

// 	target := []int{1, 2}
// 	result := findMode(tree)
// 	if !verify(target, result) {
// 		t.Errorf("expect %+v, got: %+v", target, result)
// 	}
// }

func TestFindModeIII(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}

	target := []int{1}
	result := findMode(tree)
	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}
