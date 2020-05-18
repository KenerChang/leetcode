package flattenbinarytreetolinkedlist

import (
	"testing"
)

func Verify(tree1, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}

	if tree1 == nil || tree2 == nil {
		return false
	}

	if tree1.Val != tree2.Val {
		return false
	}

	return Verify(tree1.Left, tree2.Left) && Verify(tree1.Right, tree2.Right)
}

func preorder(tree *TreeNode) []int {
	stack := []*TreeNode{}
	curNode := tree

	order := []int{}
	for len(stack) > 0 || curNode != nil {
		if curNode == nil {
			curNode = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
		}

		order = append(order, curNode.Val)

		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}
		curNode = curNode.Left
	}
	return order
}

func TestFlatten(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	target := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
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
			},
		},
	}

	flatten(tree)
	if !Verify(target, tree) {
		t.Errorf("expect: %+v, got %+v", preorder(target), preorder(tree))
	}
}
