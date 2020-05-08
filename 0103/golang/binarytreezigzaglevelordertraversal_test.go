package binarytreezigzaglevelordertraversal

import (
	"testing"
)

func verify(target, result [][]int) bool {
	if len(target) != len(result) {
		return false
	}

	size := len(target)
	for i := 0; i < size; i++ {
		targetSlice := target[i]
		resultSlice := result[i]

		if len(targetSlice) != len(resultSlice) {
			return false
		}

		for j := 0; j < len(targetSlice); j++ {
			if targetSlice[j] != resultSlice[j] {
				return false
			}
		}
	}
	return true
}

func TestZigzagLevelOrder(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	target := [][]int{
		[]int{3},
		[]int{20, 9},
		[]int{15, 7},
	}
	result := zigzagLevelOrder(tree)

	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}

func TestZigzagLevelOrderII(t *testing.T) {
	target := [][]int{}
	result := zigzagLevelOrder(nil)

	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}

func TestZigzagLevelOrderIII(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
	}
	target := [][]int{
		[]int{1},
	}
	result := zigzagLevelOrder(tree)

	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}

func TestZigzagLevelOrderIV(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	target := [][]int{
		[]int{1},
		[]int{3, 2},
		[]int{4, 5, 6, 7},
	}
	result := zigzagLevelOrder(tree)

	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}
