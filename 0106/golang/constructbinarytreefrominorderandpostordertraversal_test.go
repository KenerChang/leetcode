package constructbinarytreefrominorderandpostordertraversal

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
	postorder := []int{9, 15, 7, 20, 3}

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

	result := buildTree(inorder, postorder)
	if !verify(target, result) {
		t.Errorf("result doest not match target")
	}
}

func TestBuildTreeII(t *testing.T) {
	inorder := []int{}
	postorder := []int{}

	result := buildTree(inorder, postorder)
	if !verify(nil, result) {
		t.Errorf("result doest not match target")
	}
}

func TestBuildTreeIII(t *testing.T) {
	inorder := []int{4, 2, 5, 1, 6, 3, 7}
	postorder := []int{4, 5, 2, 6, 7, 3, 1}

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

	result := buildTree(inorder, postorder)
	if !verify(target, result) {
		t.Errorf("result doest not match target")
	}
}

func TestBuildTreeIV(t *testing.T) {
	inorder := []int{1}
	postorder := []int{1}

	target := &TreeNode{
		Val: 1,
	}

	result := buildTree(inorder, postorder)
	if !verify(target, result) {
		t.Errorf("result doest not match target")
	}
}

func TestBuildTreeV(t *testing.T) {
	inorder := []int{2, 1}
	postorder := []int{2, 1}

	target := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}

	result := buildTree(inorder, postorder)
	if !verify(target, result) {
		t.Errorf("result doest not match target")
	}
}

func TestBuildTreeVI(t *testing.T) {
	inorder := []int{1, 2}
	postorder := []int{2, 1}

	target := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}

	result := buildTree(inorder, postorder)
	if !verify(target, result) {
		t.Errorf("result doest not match target")
	}
}
