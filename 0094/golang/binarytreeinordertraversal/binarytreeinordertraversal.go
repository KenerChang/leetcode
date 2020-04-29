package binarytreeinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	return inorderTraversal(root)
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	stack := []*TreeNode{root}
	curNode := root
	traced := map[*TreeNode]bool{}

	for len(stack) > 0 {
		stackSize := len(stack)
		curNode = stack[stackSize-1]

		if _, found := traced[curNode]; found {
			result = append(result, curNode.Val)
			stack = stack[:stackSize-1]

			if curNode.Right != nil {
				stack = append(stack, curNode.Right)
			}
		} else {
			if curNode.Left != nil {
				stack = append(stack, curNode.Left)
			}
			traced[curNode] = true
		}
	}
	return result
}
