package binarytreeinordertraversal

import (
	"testing"
)

func verify(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for idx, num1 := range s1 {
		if num1 != s2[idx] {
			return false
		}
	}

	return true
}
func TestInorderTraversalI(t *testing.T) {
	input := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	result := inorderTraversal(input)
	target := []int{1, 3, 2}
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}

func TestInorderTraversalII(t *testing.T) {
	result := inorderTraversal(nil)
	target := []int{}
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}

func TestInorderTraversalIII(t *testing.T) {
	input := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	result := inorderTraversal(input)
	target := []int{1}
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}

func TestInorderTraversalIV(t *testing.T) {
	input := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: nil,
		},
		Right: nil,
	}

	result := inorderTraversal(input)
	target := []int{4, 3, 2, 1}
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}

func TestInorderTraversalV(t *testing.T) {
	input := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Left: nil,
			},
			Left: nil,
		},
		Left: nil,
	}

	result := inorderTraversal(input)
	target := []int{1, 2, 3, 4}
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}
