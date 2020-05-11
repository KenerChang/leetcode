package binarytreelevelordertraversalii

import (
	"testing"
)

func verify(levels1, levels2 [][]int) bool {
	if len(levels1) != len(levels2) {
		return false
	}

	for idx, level1 := range levels1 {
		level2 := levels2[idx]
		if len(level1) != len(level2) {
			return false
		}

		for j, num1 := range level2 {
			num2 := level2[j]
			if num1 != num2 {
				return false
			}
		}
	}
	return true
}

func TestLevelOrderBottom(t *testing.T) {
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
		[]int{15, 7},
		[]int{9, 20},
		[]int{3},
	}
	result := levelOrderBottom(tree)

	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}

func TestLevelOrderBottomII(t *testing.T) {
	target := [][]int{}
	result := levelOrderBottom(nil)

	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}

func TestLevelOrderBottomIII(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
	}

	target := [][]int{
		[]int{3},
	}
	result := levelOrderBottom(tree)

	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}

func TestLevelOrderBottomIV(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
				},
			},
		},
	}

	target := [][]int{
		[]int{6},
		[]int{5},
		[]int{4},
		[]int{3},
	}
	result := levelOrderBottom(tree)

	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}

func TestLevelOrderBottomV(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Right: &TreeNode{
			Val: 4,
			Right: &TreeNode{
				Val: 5,
				Right: &TreeNode{
					Val: 6,
				},
			},
		},
	}

	target := [][]int{
		[]int{6},
		[]int{5},
		[]int{4},
		[]int{3},
	}
	result := levelOrderBottom(tree)

	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}
