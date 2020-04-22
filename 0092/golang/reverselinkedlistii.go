package reverselinkedlistii

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == 0 {
		m++
		n++
	}
	curNode := head
	var lastNode *ListNode
	var mIdxNode *ListNode
	var preNode *ListNode
	idx := 0

	for curNode != nil {
		// put the latest node to first
		if idx >= m && idx <= n-1 {
			next := curNode.Next
			lastNode.Next = next

			curNode.Next = mIdxNode
			mIdxNode = curNode

			curNode = next

			if m != 1 {
				preNode.Next = mIdxNode
			}

		} else {
			if idx == m-2 {
				preNode = curNode
			}

			if idx == m-1 {
				lastNode = curNode
				mIdxNode = curNode
			}
			curNode = curNode.Next
		}
		idx++
	}

	if m == 1 {
		return mIdxNode
	}
	return head
}
