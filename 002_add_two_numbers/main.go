package leetcode002


type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode {}
	carry := 0
	childNode := node
	l1Val := 0
	l2Val := 0

	for l1 != nil || l2 != nil {
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		} else {
			l1Val = 0
		}

		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		} else {
			l2Val = 0
		}

		addNumber := l1Val + l2Val + carry
		carry = addNumber / 10
		childNode.Next = &ListNode { Val: addNumber % 10 }
		childNode = childNode.Next
	}

	if carry == 1 {
		childNode.Next = &ListNode { Val: 1 }
		childNode = childNode.Next
	}
	
	return node.Next
}
