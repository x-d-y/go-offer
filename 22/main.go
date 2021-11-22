package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	h := head
	var l *ListNode

	for i := 1; head != nil; i++ {
		head = head.Next
		if i > k {
			l = l.Next
		} else {
			l = h
		}
	}
	return l
}
