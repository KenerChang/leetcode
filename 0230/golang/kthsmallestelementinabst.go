package kthsmallestelementinabst

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func kthSmallest(root *TreeNode, k int) int {
	results := kthSmallestRecursive(root, k, []int{})
	return results[k-1]
}

func kthSmallestRecursive(root *TreeNode, k int, track []int) []int {
	if len(track) == k || root == nil {
		return track
	}

	track = kthSmallestRecursive(root.Left, k, track)
	track = append(track, root.Val)
	track = kthSmallestRecursive(root.Right, k, track)
	return track
}
