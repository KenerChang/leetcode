package convertsortedlisttobinarysearchtree

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

func createList(nums []int) *ListNode {
	var first, curNode *ListNode
	for idx, num := range nums {
		if idx == 0 {
			first = &ListNode{
				Val: num,
			}
			curNode = first
		} else {
			node := &ListNode{
				Val: num,
			}
			curNode.Next = node
			curNode = node
		}
	}
	return first
}

func TestSortedListToBST(t *testing.T) {
	list := createList([]int{-10, -3, 0, 5, 9})
	target := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: -3,
			Left: &TreeNode{
				Val: -10,
			},
		},
		Right: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 5,
			},
		},
	}

	result := sortedListToBST(list)
	if !verify(target, result) {
		t.Errorf("result does not match target")
	}
}

func TestSortedListToBSTII(t *testing.T) {
	result := sortedListToBST(nil)
	if !verify(nil, result) {
		t.Errorf("result does not match target")
	}
}

func TestSortedListToBSTIII(t *testing.T) {
	list := createList([]int{0})
	target := &TreeNode{
		Val: 0,
	}

	result := sortedListToBST(list)
	if !verify(target, result) {
		t.Errorf("result does not match target")
	}
}

func TestSortedListToBSTIV(t *testing.T) {
	list := createList([]int{0, 1})
	target := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
	}

	result := sortedListToBST(list)
	if !verify(target, result) {
		t.Errorf("result does not match target")
	}
}

func TestSortedArrayToBSTV(t *testing.T) {
	list := createList([]int{0, 1, 2})
	target := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	result := sortedListToBST(list)
	if !verify(target, result) {
		t.Errorf("result does not match target")
	}
}
