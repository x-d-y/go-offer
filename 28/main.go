package main

import (
	"fmt"
	"strconv"
)

func main() {
	root := buildTree([]string{"1", "2", "2", "null", "3", "null", "3"})

	fmt.Println(isSymmetric(root))

}

func isSymmetric(root *TreeNode) bool {
	return re(root, root)
}

func re(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	if l.Val == r.Val {
		return re(l.Left, r.Right) && re(l.Right, r.Left)
	}
	return false
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
