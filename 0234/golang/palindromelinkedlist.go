package palindromelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	oneStep := head
	twoStep := head

	nums := []int{}
	for twoStep != nil {
		twoStep = twoStep.Next
		if twoStep == nil {
			// reach end, and node number is odd
			oneStep = oneStep.Next
			break
		}

		twoStep = twoStep.Next
		if twoStep == nil {
			// reach end, and node number is even
			nums = append(nums, oneStep.Val)
			oneStep = oneStep.Next
			break
		}

		nums = append(nums, oneStep.Val)
		oneStep = oneStep.Next
	}

	targetPos := len(nums) - 1
	for oneStep != nil {
		if oneStep.Val != nums[targetPos] {
			return false
		}

		oneStep = oneStep.Next
		targetPos--
	}
	return true
}
