package leetcode002

import(
	"testing"
)

func makeListNode(e []int) *ListNode {
	node := &ListNode{ Val: e[0] }
	childNode := node
	for i := 1; i < len(e); i++ {
		childNode.Next = &ListNode{ Val: e[i] }
		childNode = childNode.Next
	}
	return node
}

func assertListNode(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil || l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}


func Test_Example_1(t *testing.T) {
	l1 := makeListNode([]int{2,4,3})
	l2 := makeListNode([]int{5,6,4})
	actual := addTwoNumbers(l1, l2)
	expected := makeListNode([]int{7,0,8})

	if assertListNode(actual, expected) == false {
		t.Errorf("got: %v, actual: %v", actual, expected)
	}
}

func Test_Example_2(t *testing.T) {
	l1 := makeListNode([]int{5})
	l2 := makeListNode([]int{5})
	actual := addTwoNumbers(l1, l2)
	expected := makeListNode([]int{0,1})

	if assertListNode(actual, expected) == false {
		t.Errorf("got: %v, actual: %v", actual, expected)
	}
}

