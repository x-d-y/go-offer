package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	lens := len(res)
	for i := 0; i <= lens/2; i++ {
		res[i], res[lens-i-1] = res[lens-i-1], res[i]
	}
	return res
}

func appendLeft(num int, nums []int) []int {
	return append([]int{num}, nums...)
}
