package main

import (
	"fmt"
	"strconv"
)

func main() {
	root := buildTree([]string{"4", "2", "7", "1", "3", "6", "9"})
	res := mirrorTree(root)
	fmt.Println(res.Left.Left.Val)
}

func mirrorTree(root *TreeNode) *TreeNode {
	r(root)
	return root
}

func r(root *TreeNode) {
	if root == nil{
		return
	}
	root.Left, root.Right = root.Right, root.Left
	r(root.Left)
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
