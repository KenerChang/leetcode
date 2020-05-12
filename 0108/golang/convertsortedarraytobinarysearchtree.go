package convertsortedarraytobinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}

	if len(nums) == 2 {
		return &TreeNode{
			Val: nums[1],
			Left: &TreeNode{
				Val: nums[0],
			},
		}
	}

	if len(nums) == 3 {
		return &TreeNode{
			Val: nums[1],
			Left: &TreeNode{
				Val: nums[0],
			},
			Right: &TreeNode{
				Val: nums[2],
			},
		}
	}

	rootIdx := len(nums) / 2

	tree := &TreeNode{
		Val:   nums[rootIdx],
		Left:  sortedArrayToBST(nums[0:rootIdx]),
		Right: sortedArrayToBST(nums[rootIdx+1:]),
	}

	return tree
}
