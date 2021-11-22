package main

func main() {
	headA := &ListNode{Val: 2}
	headB := &ListNode{Val: 3}
	headA.Next = headB
	getIntersectionNode(headA, headB)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		}
		if b == nil {
			b = headA
		}
		if a == b {
			break
		}
		a = a.Next
		b = b.Next
	}
	return a
}
