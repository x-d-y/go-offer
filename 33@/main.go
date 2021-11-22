package main

import "fmt"

func main() {
	fmt.Println(verifyPostorder([]int{1,3,2,5}))
}

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}
	var lflag int
	for i := len(postorder) - 2; i >= 0; i-- {
		if postorder[i] > postorder[len(postorder)-1] {
			if lflag > 0 {
				return false
			}

		}
		if postorder[i] < postorder[len(postorder)-1] {
			lflag++
		}
	}
	return verifyPostorder(postorder[:lflag]) && verifyPostorder(postorder[lflag:len(postorder)-1])
}
