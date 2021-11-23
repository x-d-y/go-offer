package main

import (
	"fmt"
	"strconv"
)

// func main() {
// 	root := buildTree([]string{"5", "3", "6", "2", "4", "null", "null", "1"})
// 	fmt.Println(kthLargest(root, 6).Val)

// }

// var i int

// func kthLargest(root *TreeNode, k int) *TreeNode {
// 	i = 0
// 	return r(root, k)
// }

// func r(root *TreeNode, k int) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}

// 	if left := r(root.Left, k); left != nil {
// 		return left
// 	}

// 	i++
// 	if i == k {
// 		return root
// 	}

// 	if right := r(root.Right, k); right != nil {
// 		return right
// 	}
// 	return nil

// }

var res []int

func kthLargest(root *TreeNode, k int) int {
	res = make([]int, 0)
	r(root)
	return res[len(res)-k]
}

func r(root *TreeNode) {
	if root == nil {
		return
	}
	r(root.Left)
	res = append(res, root.Val)
	r(root.Right)
}

func buildTree(nums []string) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := new(TreeNode)
	n, _ := strconv.Atoi(nums[0])
	root.Val = n
	temp := make([]*TreeNode, 0)
	temp = append(temp, root)
	start := 1
	for i := 0; i < len(temp); i++ {
		if temp[i] == nil {
			continue
		}
		l, r := false, false
		for start < len(nums) {
			var cNode *TreeNode
			if nums[start] != "null" {
				cNode = new(TreeNode)
				n, _ := strconv.Atoi(nums[start])
				cNode.Val = n
			}
			temp = append(temp, cNode)
			start++
			if !l {
				temp[i].Left = cNode
				l = true
				continue
			}
			if !r {
				temp[i].Right = cNode
				break
			}
		}
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
