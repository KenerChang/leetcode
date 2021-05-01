package binarytreepaths

import (
	"sort"
	"testing"
)

func verify(target, result []string) bool {
	if len(target) != len(result) {
		return false
	}

	sort.Strings(target)
	sort.Strings(result)

	for i := 0; i < len(target); i++ {
		if target[i] != result[i] {
			return false
		}
	}
	return true
}

func TestBinaryTreePaths(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	result := binaryTreePaths(root)

	target := []string{"1->2->5", "1->3"}
	if !verify(target, result) {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestBinaryTreePathsII(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}

	result := binaryTreePaths(root)

	target := []string{"1"}
	if !verify(target, result) {
		t.Errorf("expect %s, got %s", target, result)
	}
}
