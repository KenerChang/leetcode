package recoverbinarysearchtree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}

	cur := root
	stack := []*TreeNode{}
	var prev, node1, node2 *TreeNode
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]

			if prev != nil && cur.Val < prev.Val {
				node2 = cur

				if node1 != nil {
					break
				} else {
					node1 = prev
				}
			}

			prev = cur
			cur = cur.Right
		}
	}

	node1.Val, node2.Val = node2.Val, node1.Val
}
