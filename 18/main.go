package main

func main() {

}

func deleteNode(head *ListNode, val int) *ListNode {
	h := head
	var pre *ListNode
	for head != nil {
		if head.Val == val {
			if pre != nil {
				pre.Next = head.Next
			} else {
				return head.Next
			}
		}
		pre = head
		head = head.Next
	}
	return h
}

type ListNode struct {
	Val  int
	Next *ListNode
}
