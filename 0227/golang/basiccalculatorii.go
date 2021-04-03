package basiccalculatorii

import (
	"container/list"
	"strconv"
)

const (
	TypeOperand  = 1
	TypeOperator = 2
)

type Node struct {
	Type     int
	Val      int
	Operator rune
}

var (
	operatorsMap map[rune]int = map[rune]int{
		'/': 2,
		'*': 2,
		'+': 1,
		'-': 1,
	}
	operandsMap map[rune]bool = map[rune]bool{
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}
)

func calculate(s string) int {
	queue := list.New()

	operand := []rune{}
	var lastOperator *list.Element
	var prevOperatorNode *Node
	for _, char := range s {
		if _, ok := operandsMap[char]; ok {
			operand = append(operand, char)
		}

		if _, ok := operatorsMap[char]; ok {
			value, _ := strconv.Atoi(string(operand))
			operandNode := &Node{
				Type: TypeOperand,
				Val:  value,
			}
			operand = []rune{}

			operatorNode := &Node{
				Type:     TypeOperator,
				Operator: char,
			}

			if prevOperatorNode != nil {
				if lastOperator != nil {
					lastOperatorNode := lastOperator.Value.(*Node)
					lastOperatorPriority := operatorsMap[lastOperatorNode.Operator]

					prevOperatorPriority := operatorsMap[prevOperatorNode.Operator]

					if prevOperatorPriority > lastOperatorPriority {
						queue.InsertBefore(operandNode, lastOperator)
						queue.InsertBefore(prevOperatorNode, lastOperator)
					} else {
						queue.PushBack(operandNode)
						lastOperator = queue.PushBack(prevOperatorNode)
					}
					prevOperatorNode = operatorNode

				} else {
					queue.PushBack(operandNode)
					lastOperator = queue.PushBack(prevOperatorNode)
					prevOperatorNode = operatorNode
				}
			} else {
				queue.PushBack(operandNode)
				prevOperatorNode = operatorNode
				// lastOperator = queue.PushBack(operatorNode)
			}
		}
	}

	value, _ := strconv.Atoi(string(operand))
	if prevOperatorNode == nil {
		return value
	}

	operandNode := &Node{
		Val:  value,
		Type: TypeOperand,
	}

	if lastOperator == nil {
		queue.PushBack(operandNode)
		queue.PushBack(prevOperatorNode)
	} else {
		lastOperatorNode := lastOperator.Value.(*Node)
		lastOperatorPriority := operatorsMap[lastOperatorNode.Operator]

		prevOperatorPriority := operatorsMap[prevOperatorNode.Operator]

		if prevOperatorPriority > lastOperatorPriority {
			queue.InsertBefore(operandNode, lastOperator)
			queue.InsertBefore(prevOperatorNode, lastOperator)
		} else {
			queue.PushBack(operandNode)
			queue.PushBack(prevOperatorNode)
		}
	}

	var leftOperand, rightOperand *Node
	var leftElement, rightElement *list.Element
	for e := queue.Front(); e != nil; {
		node := e.Value.(*Node)
		// switch node.Type {
		// case TypeOperator:
		// 	fmt.Printf("operator: %s\n", string(node.Operator))
		// case TypeOperand:
		// 	fmt.Printf("operand: %d\n", node.Val)
		// }

		if node.Type == TypeOperator {
			rightElement = e.Prev()
			rightOperand = rightElement.Value.(*Node)

			if leftElement == nil {
				leftElement = e.Prev().Prev()
				leftOperand = leftElement.Value.(*Node)
			}

			// fmt.Printf("left operand: %d, right operand: %d\n", leftOperand.Val, rightOperand.Val)

			switch node.Operator {
			case '+':
				rightOperand.Val = leftOperand.Val + rightOperand.Val
			case '-':
				rightOperand.Val = leftOperand.Val - rightOperand.Val
			case '*':
				rightOperand.Val = leftOperand.Val * rightOperand.Val
			case '/':
				rightOperand.Val = leftOperand.Val / rightOperand.Val
			}

			queue.Remove(rightElement)
			queue.Remove(leftElement)
			next := e.Next()
			queue.Remove(e)
			e = next

			if e != nil {
				queue.InsertBefore(rightOperand, e)
			}

			leftElement = nil
		} else {
			e = e.Next()
		}

	}

	return rightOperand.Val
}
