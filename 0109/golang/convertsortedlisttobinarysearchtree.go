package convertsortedlisttobinarysearchtree

// import (
// 	"fmt"
// )

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	curNode := head
	nums := []int{}
	for curNode != nil {
		nums = append(nums, curNode.Val)
		curNode = curNode.Next
	}
	// fmt.Printf("nums: %+v\n", nums)
	return sortedListToBSTImpl(nums)
}

func sortedListToBSTImpl(nums []int) *TreeNode {
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
		Left:  sortedListToBSTImpl(nums[0:rootIdx]),
		Right: sortedListToBSTImpl(nums[rootIdx+1:]),
	}

	return tree
}
