package inordersuccessorinbst

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

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	// if a node.Val is greater than p.Val, then this node is
	// impossible to be the successor of p. So we can keep traversed
	// node in a stack and only need to keep node which Val is greater
	// then p.Val in the stack, and the node on top of the stack is
	// the successor of p.

	var successor *TreeNode

	cur := root
	for cur.Val != p.Val {
		if cur.Val > p.Val {
			// go left & keep the node
			successor = cur

			cur = cur.Left
		} else {
			// go right
			cur = cur.Right
		}
	}

	cur = cur.Right
	for cur != nil {
		successor = cur
		cur = cur.Left
	}

	return successor
}
