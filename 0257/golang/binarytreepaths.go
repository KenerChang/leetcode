package binarytreepaths

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	return binaryTreePathsRecursive(root, []string{}, "")
}

func binaryTreePathsRecursive(root *TreeNode, results []string, path string) []string {
	if path == "" {
		path = fmt.Sprint(root.Val)
	} else {
		path = fmt.Sprintf("%s->%d", path, root.Val)
	}

	if root.Left == nil && root.Right == nil {
		// a leaf
		results = append(results, path)
		return results
	}

	if root.Left != nil {
		results = binaryTreePathsRecursive(root.Left, results, path)
	}

	if root.Right != nil {
		results = binaryTreePathsRecursive(root.Right, results, path)
	}

	return results
}
