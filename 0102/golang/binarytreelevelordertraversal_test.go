package binarytreelevelordertraversal

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

func TestLevelOrderI(t *testing.T) {
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
		[]int{9, 20},
		[]int{15, 7},
	}
	result := levelOrder(tree)
	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}

func TestLevelOrderII(t *testing.T) {
	target := [][]int{}

	result := levelOrder(nil)
	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}

func TestLevelOrderIII(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
	}
	target := [][]int{
		[]int{5},
	}

	result := levelOrder(tree)
	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}

func TestLevelOrderIV(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 7,
		},
	}
	target := [][]int{
		[]int{5},
		[]int{6, 7},
	}

	result := levelOrder(tree)
	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}

func TestLevelOrderV(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 7,
						},
					},
				},
			},
		},
	}
	target := [][]int{
		[]int{5},
		[]int{6},
		[]int{7},
		[]int{7},
		[]int{7},
		[]int{7},
	}

	result := levelOrder(tree)
	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}

func TestLevelOrderVI(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 7,
				Right: &TreeNode{
					Val: 7,
					Right: &TreeNode{
						Val: 7,
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
		},
	}
	target := [][]int{
		[]int{5},
		[]int{6},
		[]int{7},
		[]int{7},
		[]int{7},
		[]int{7},
	}

	result := levelOrder(tree)
	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}
