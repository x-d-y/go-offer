package main

func main() {

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}
	l1.Next = mergeTwoLists(l1.Next, l2)
	return l1
}

type ListNode struct {
	Val  int
	Next *ListNode
}
