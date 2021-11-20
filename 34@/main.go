package main

import (
	"fmt"
	"strconv"
)

func main() {
	root := buildTree([]string{"5", "4", "8", "11", "null", "13", "4", "7", "2", "null", "null", "5", "1"})
	fmt.Println(pathSum(root, 22))
}

var res [][]int
var subRes []int

func pathSum(root *TreeNode, target int) [][]int {
	res = make([][]int, 0)
	subRes = make([]int, 0)
	r(root, target)
	return res
}

func r(root *TreeNode, target int) {
	if root == nil {
		subRes = append(subRes, 0)
		return
	}
	subRes = append(subRes, root.Val)
	if root.Val == target && root.Left == nil && root.Right == nil {
		res = append(res, append(make([]int, 0), subRes...))
		return
	}
	r(root.Left, target-root.Val)
	subRes = subRes[:len(subRes)-1]
	r(root.Right, target-root.Val)
	subRes = subRes[:len(subRes)-1]
	return
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
