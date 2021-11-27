package main

import (
	"fmt"
	"strconv"
)

func main() {
	root := buildTree([]string{"1", "2", "2", "3", "3", "null", "null", "4", "4"})
	fmt.Println(isBalanced(root))
}

func isBalanced(root *TreeNode) bool {
	is = true
	r(root)
	return is
}

var is bool

func r(root *TreeNode) int {
	if !is {
		return 0
	}
	if root == nil {
		return 0
	}
	deepl := r(root.Left)
	deepr := r(root.Right)
	s := deepl - deepr

	if s <= 1 && s >= -1 {
		return max(deepl, deepr) + 1
	} else {
		is = false
		return 0
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
