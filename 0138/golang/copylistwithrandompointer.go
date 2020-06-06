package copylistwithrandompointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	mapping := map[*Node]*Node{}
	curHead := head
	var copied *Node
	var copiedHead *Node
	for curHead != nil {
		copied = mapping[curHead]
		if copied == nil {
			copied = &Node{
				Val: curHead.Val,
			}

			mapping[curHead] = copied
		}

		if curHead.Next != nil {
			if next, found := mapping[curHead.Next]; found {
				copied.Next = next
			} else {
				next := &Node{
					Val: curHead.Next.Val,
				}
				copied.Next = next
				mapping[curHead.Next] = next
			}
		}

		if curHead.Random != nil {
			if random, found := mapping[curHead.Random]; found {
				copied.Random = random
			} else {
				random := &Node{
					Val: curHead.Random.Val,
				}
				copied.Random = random
				mapping[curHead.Random] = random
			}
		}

		if copiedHead == nil {
			copiedHead = copied
		}

		curHead = curHead.Next
	}
	return copiedHead
}
