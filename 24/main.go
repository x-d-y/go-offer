package mian

func main() {

}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var next *ListNode
	var pre *ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

type ListNode struct {
	Val  int
	Next *ListNode
}
