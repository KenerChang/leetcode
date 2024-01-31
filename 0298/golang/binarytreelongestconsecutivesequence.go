package binarytreelongestconsecutivesequence

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestConsecutive(root *TreeNode) int {
	return longestConsecutiveRecursive(root, 1)
}

func longestConsecutiveRecursive(root *TreeNode, sequenceCount int) int {
	var longestFromLeft int
	if root.Left != nil {
		leftSequenceCount := 1
		if root.Left.Val == root.Val+1 {
			leftSequenceCount = sequenceCount + 1

		}
		longestFromLeft = longestConsecutiveRecursive(root.Left, leftSequenceCount)
	}

	var longestFromRight int
	if root.Right != nil {
		rightSequenceCount := 1
		if root.Right.Val == root.Val+1 {
			rightSequenceCount = sequenceCount + 1
		}
		longestFromRight = longestConsecutiveRecursive(root.Right, rightSequenceCount)
	}

	return max(sequenceCount, longestFromLeft, longestFromRight)
}

func max(nums ...int) int {
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > result {
			result = nums[i]
		}
	}

	return result
}
