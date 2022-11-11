package stepbystepdirectionsfromabinarytreenodetoanother

import "strings"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findPath(root *TreeNode, target int) (string, bool) {
	if root == nil {
		return "", false
	}

	if root.Val == target {
		return "", true
	}

	lPath, ok := findPath(root.Left, target)
	if ok {
		return "L" + lPath, ok
	}

	rPath, ok := findPath(root.Right, target)
	if ok {
		return "R" + rPath, ok
	}

	return "", false
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	// To solve this problem, we need a util function to find the path from root to target

	// first find path from root to startValue
	pathOfStartValue, _ := findPath(root, startValue)

	// second find path from root to destValue
	pathOfDestValue, _ := findPath(root, destValue)

	// remove prefix pathes
	var sp, dp int
	for sp < len(pathOfStartValue) && dp < len(pathOfDestValue) {
		if pathOfStartValue[sp] != pathOfDestValue[dp] {
			break
		}

		sp++
		dp++
	}

	var result string
	if sp < len(pathOfStartValue) {
		result = strings.Repeat("U", len(pathOfStartValue)-sp)
	}

	if dp < len(pathOfDestValue) {
		result += pathOfDestValue[dp:]
	}

	return result
}
