package constructbinarytreefrompreorderandinordertraversal

import (
	"testing"
)

func verify(tree1, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}

	if tree1 == nil || tree2 == nil {
		return false
	}

	if tree1.Val != tree2.Val {
		return false
	}

	return verify(tree1.Left, tree2.Left) && verify(tree1.Right, tree2.Right)
}

func TestBuildTree(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	preorder := []int{3, 9, 20, 15, 7}

	target := &TreeNode{
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
	result := buildTree(preorder, inorder)

	if !verify(target, result) {
		t.Errorf("target does not match result")
	}
}

func TestBuildTreeII(t *testing.T) {
	preorder := []int{1, 2, 4, 5, 3, 6, 7}
	inorder := []int{4, 2, 5, 1, 6, 3, 7}

	target := &TreeNode{
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
	result := buildTree(preorder, inorder)

	if !verify(target, result) {
		t.Errorf("target does not match result")
	}
}

func TestBuildTreeIII(t *testing.T) {
	preorder := []int{}
	inorder := []int{}

	result := buildTree(preorder, inorder)

	if !verify(nil, result) {
		t.Errorf("target does not match result")
	}
}

func TestBuildTreeIV(t *testing.T) {
	preorder := []int{1}
	inorder := []int{1}
	target := &TreeNode{
		Val: 1,
	}

	result := buildTree(preorder, inorder)

	if !verify(target, result) {
		t.Errorf("target does not match result")
	}
}

func TestBuildTreeV(t *testing.T) {
	preorder := []int{1, 2}
	inorder := []int{2, 1}
	target := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}

	result := buildTree(preorder, inorder)

	if !verify(target, result) {
		t.Errorf("target does not match result")
	}
}

func TestBuildTreeVI(t *testing.T) {
	preorder := []int{1, 2}
	inorder := []int{1, 2}
	target := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}

	result := buildTree(preorder, inorder)

	if !verify(target, result) {
		t.Errorf("target does not match result")
	}
}
